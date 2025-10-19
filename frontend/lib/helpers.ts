import { z } from "zod";
import { Platform, FieldKey, REQUIRED_BY_PLATFORM, FormData } from "./types";

/**
 * Figure out what fields the user needs to fill out based on their platform choices
 * 
 * When someone selects YouTube and Instagram, this tells us they need to fill out
 * fields like "title" (for YouTube), "caption" (for Instagram), etc.
 * We combine all the required fields and remove any duplicates.
 */
export function getRequiredFields(selected: Platform[]): FieldKey[] {
  // Get all required fields from each selected platform
  const allFields = selected.flatMap(platform => REQUIRED_BY_PLATFORM[platform]);
  // Remove duplicates - if both YouTube and Pinterest need "title", we only show it once
  return Array.from(new Set(allFields));
}

/**
 * Get a breakdown of which fields are needed for each specific platform
 * 
 * This is useful when we want to show the user "YouTube needs these fields, 
 * Instagram needs these other fields" instead of just a combined list.
 */
export function getFieldsByPlatform(selected: Platform[]): Record<Platform, FieldKey[]> {
  const result: Record<Platform, FieldKey[]> = {} as Record<Platform, FieldKey[]>;
  
  // For each platform the user selected, get its specific required fields
  selected.forEach(platform => {
    result[platform] = REQUIRED_BY_PLATFORM[platform];
  });
  
  return result;
}

/**
 * Check if a specific field is needed for any of the platforms the user selected
 * 
 * Example: If user selected YouTube and Instagram, and we ask "is 'title' required?",
 * this returns true because YouTube needs it (even though Instagram doesn't).
 */
export function isFieldRequired(field: FieldKey, selected: Platform[]): boolean {
  return selected.some(platform => REQUIRED_BY_PLATFORM[platform].includes(field));
}

/**
 * Find out which of the user's selected platforms actually need a specific field
 * 
 * Example: If user selected YouTube, Instagram, and LinkedIn, and we ask "which platforms 
 * need the 'title' field?", this would return ["youtube"] because only YouTube requires it.
 */
export function getPlatformsForField(field: FieldKey, selected: Platform[]): Platform[] {
  return selected.filter(platform => REQUIRED_BY_PLATFORM[platform].includes(field));
}

/**
 * Create form validation rules that change based on what platforms the user selected
 * 
 * This is like creating a smart form that knows "if they're posting to YouTube, 
 * they need a title and video. If they're posting to Instagram, they need a caption 
 * and either an image or video." Each field gets different validation rules.
 */
export function buildZodSchema(fields: FieldKey[], selectedPlatforms: Platform[]) {
  const schemaFields: Record<string, z.ZodTypeAny> = {};

  // Go through each field that needs validation and set up the right rules
  fields.forEach(field => {
    const platforms = getPlatformsForField(field, selectedPlatforms);
    
    switch (field) {
      // These are simple text fields that just need to not be empty
      case "title":
      case "description":
      case "caption":
      case "subreddit":
      case "board_id":
      case "author":
        schemaFields[field] = z.string().min(1, `${field} is required`);
        break;
        
      // These need to be valid URLs (for images, videos, links)
      case "image_url":
      case "thumbnail":
      case "video_file":
      case "link":
        schemaFields[field] = z.string().url(`Valid ${field} URL is required`);
        break;
        
      // Tags come in as a comma-separated string, but we want an array
      case "tags":
        schemaFields[field] = z.string().optional().transform((val) => 
          val ? val.split(',').map(tag => tag.trim()).filter(Boolean) : []
        );
        break;
        
      // Reddit posts need content if it's a "self" post (text post, not a link)
      case "content":
        if (platforms.includes("reddit")) {
          schemaFields[field] = z.string().min(1, "Content is required for Reddit self posts");
        } else {
          schemaFields[field] = z.string().optional();
        }
        break;
        
      // YouTube privacy settings - user picks from a dropdown
      case "privacy_status":
        schemaFields[field] = z.enum(["public", "unlisted", "private"]).optional();
        break;
        
      // Reddit post types - what kind of post is this?
      case "post_type":
        schemaFields[field] = z.enum(["self", "link", "image"]);
        break;
        
      // LinkedIn visibility - who can see this post?
      case "visibility":
        schemaFields[field] = z.enum(["PUBLIC", "CONNECTIONS"]);
        break;
        
      // Reddit NSFW flag - is this content not safe for work?
      case "nsfw":
        schemaFields[field] = z.boolean().optional();
        break;
        
      // These are optional text fields for various platform-specific settings
      case "category_id":
      case "location_id":
      case "source_type":
      case "lifecycle_state":
        schemaFields[field] = z.string().optional();
        break;
        
      // Instagram user tags - mention other users (also comma-separated)
      case "user_tags":
        schemaFields[field] = z.string().optional().transform((val) => 
          val ? val.split(',').map(tag => tag.trim()).filter(Boolean) : []
        );
        break;
        
      // Fallback for any field we didn't specifically handle
      default:
        schemaFields[field] = z.string().optional();
    }
  });

  // Instagram has a special rule: you MUST have either an image OR a video
  // This is different from other platforms that might be more flexible
  if (selectedPlatforms.includes("instagram")) {
    const instagramSchema = z.object({
      image_url: z.string().optional(),
      video_file: z.string().optional(),
    }).refine(
      (data) => data.image_url || data.video_file,
      {
        message: "Instagram requires either image_url or video_file",
        path: ["image_url"]
      }
    );
    
    // Add Instagram's special validation to our main validation rules
    Object.assign(schemaFields, instagramSchema.shape);
  }

  // Pinterest requires an image (not video) - must be image format
  if (selectedPlatforms.includes("pinterest")) {
    const pinterestSchema = z.object({
      image_url: z.string().min(1, "Pinterest requires an image"),
      video_file: z.string().optional(),
    }).refine(
      (data) => {
        // Check if image_url points to an image file (not video)
        if (!data.image_url) return false;
        const imageExtensions = ['.jpg', '.jpeg', '.png', '.gif', '.webp', '.bmp', '.tiff'];
        const hasImageExtension = imageExtensions.some(ext => 
          data.image_url.toLowerCase().includes(ext)
        );
        return hasImageExtension && !data.video_file;
      },
      {
        message: "Pinterest only supports image formats (JPEG, PNG, GIF, WEBP, BMP, TIFF). Videos are not allowed.",
        path: ["image_url"]
      }
    );
    
    // Add Pinterest's special validation to our main validation rules
    Object.assign(schemaFields, pinterestSchema.shape);
  }

  return z.object(schemaFields);
}

/**
 * Take the form data the user filled out and create a flat JSON payload for the backend
 * 
 * This creates a single flat structure with all fields. Unselected platforms will have
 * empty/undefined values, making it much easier for the backend to parse.
 */
export function transformToSubmitData(formData: FormData, selectedPlatforms: Platform[]): any {
  // Create a flat structure with all possible fields
  // The backend can easily check which platforms are selected and use the relevant fields
  return {
    platforms: selectedPlatforms,        // Which platforms to post to
    // All fields in one flat structure - unselected platforms will have empty values
    title: formData.title || "",
    description: formData.description || "",
    caption: formData.caption || "",
    tags: formData.tags || [],
    image_url: formData.image_url || "",
    thumbnail: formData.thumbnail || "",
    video_file: formData.video_file || "",
    link: formData.link || "",
    content: formData.content || "",
    
    // Platform-specific fields
    category_id: formData.category_id || "",
    privacy_status: formData.privacy_status || "",
    location_id: formData.location_id || "",
    user_tags: formData.user_tags || '',
    subreddit: formData.subreddit || "",
    post_type: formData.post_type || "",
    nsfw: formData.nsfw || false,
    board_id: formData.board_id || "",
    source_type: formData.source_type || "",
    author: formData.author || "",
    lifecycle_state: formData.lifecycle_state || "",
    visibility: formData.visibility || ""
  };
}
