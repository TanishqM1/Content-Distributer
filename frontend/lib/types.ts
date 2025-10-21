export type Platform = "youtube" | "reddit" | "instagram" | "pinterest" | "linkedin";

// canonical field keys used across the UI + JSON output
export type FieldKey =
  | "title" | "description" | "caption" | "tags"
  | "image_url" | "thumbnail" | "video_file" | "link"
  | "content" // reddit self-post body (aka text)
  | "category_id" | "privacy_status"            // YouTube
  | "location_id" | "user_tags"                 // Instagram
  | "subreddit" | "post_type" | "nsfw"          // Reddit
  | "board_id" | "source_type"                  // Pinterest
  | "author" | "lifecycle_state" | "visibility" // LinkedIn
  ;

export const REQUIRED_BY_PLATFORM: Record<Platform, FieldKey[]> = {
  youtube:    ["title", "description", "video_file", "category_id", "privacy_status"], // plus optional: tags
  reddit:     ["title", "content", "subreddit", "post_type", "nsfw"],      // if link post: use `link` instead of `content`
  instagram:  ["caption", "image_url", "user_tags"],  // or video_file; at least one media field required
  pinterest:  ["title", "image_url", "board_id", "description", "link", "source_type"],
  linkedin:   ["author", "visibility", "caption", "lifecycle_state"],    // caption used as post text
};

export const PLATFORM_NAMES: Record<Platform, string> = {
  youtube: "YouTube", 
  reddit: "Reddit",
  instagram: "Instagram",
  pinterest: "Pinterest",
  linkedin: "LinkedIn"
};

export const FIELD_DESCRIPTIONS: Record<FieldKey, string> = {
  title: "Title of your post/video (required)",
  description: "Detailed description of your content (required)",
  caption: "Short caption for Instagram/LinkedIn posts",
  tags: "Keywords to help people find your content (comma-separated)",
  image_url: "URL of the image you want to post",
  thumbnail: "Thumbnail image URL",
  video_file: "Video file (handled by drag & drop above)",
  link: "External URL to share",
  content: "Text content for Reddit posts",
  category_id: "YouTube category (required for YouTube)",
  privacy_status: "Who can see your video (required for YouTube)",
  location_id: "Instagram location (optional)",
  user_tags: "Instagram users to tag (comma-separated usernames)",
  subreddit: "Which subreddit to post to (required for Reddit)",
  post_type: "Type of Reddit post (required for Reddit)",
  nsfw: "Is this content NSFW? (required for Reddit)",
  board_id: "Pinterest board to pin to (required for Pinterest)",
  source_type: "Type of media source (required for Pinterest)",
  author: "LinkedIn author URN (required for LinkedIn)",
  lifecycle_state: "Publication state (required for LinkedIn)",
  visibility: "Who can see your post (required for LinkedIn)"
};

export const FIELD_APPLIES_TO: Record<FieldKey, Platform[]> = {
  title: ["youtube", "reddit", "pinterest"],
  description: ["youtube", "pinterest", "linkedin"],
  caption: ["instagram", "linkedin"],
  tags: ["youtube"],
  image_url: ["instagram", "pinterest", "reddit", "linkedin"],
  thumbnail: [],
  video_file: ["youtube", "linkedin", "instagram"],
  link: ["reddit", "pinterest"],
  content: ["reddit"],
  category_id: ["youtube"],
  privacy_status: ["youtube"],
  location_id: ["instagram"],
  user_tags: ["instagram"],
  subreddit: ["reddit"],
  post_type: ["reddit"],
  nsfw: ["reddit"],
  board_id: ["pinterest"],
  source_type: ["pinterest"],
  author: ["linkedin"],
  lifecycle_state: ["linkedin"],
  visibility: ["linkedin"]
};

export interface FormData {
  // Common fields
  title?: string;
  description?: string;
  caption?: string;
  tags?: string[];
  image_url?: string;
  thumbnail?: string;
  video_file?: string;
  link?: string;
  content?: string;
  
  // Platform-specific fields
  category_id?: string;
  privacy_status?: "public" | "unlisted" | "private";
  location_id?: string;
  user_tags?: string[];
  subreddit?: string;
  post_type?: "self" | "link" | "image";
  nsfw?: boolean;
  board_id?: string;
  source_type?: string;
  author?: string;
  lifecycle_state?: string;
  visibility?: "PUBLIC" | "CONNECTIONS";
}

export interface SubmitData {
  platforms: Platform[];
  // All fields in one flat structure - unselected platforms will have empty values
  title?: string;
  description?: string;
  caption?: string;
  tags?: string[];
  image_url?: string;
  thumbnail?: string;
  video_file?: string;
  link?: string;
  content?: string;
  
  // Platform-specific fields
  category_id?: string;
  privacy_status?: "public" | "unlisted" | "private";
  location_id?: string;
  user_tags?: string[];
  subreddit?: string;
  post_type?: "self" | "link" | "image";
  nsfw?: boolean;
  board_id?: string;
  source_type?: string;
  author?: string;
  lifecycle_state?: string;
  visibility?: "PUBLIC" | "CONNECTIONS";
}
