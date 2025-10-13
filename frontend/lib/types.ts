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
  instagram:  ["caption", "image_url"],               // or video_file; at least one media field required
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
  title: "Title of post/video",
  description: "Long description/body",
  caption: "Short caption / post text",
  tags: "Keywords array",
  image_url: "Image URL",
  thumbnail: "Thumbnail image",
  video_file: "Video file URL/path",
  link: "External URL",
  content: "Text body (self post)",
  category_id: "YouTube category",
  privacy_status: "Visibility",
  location_id: "Location tag",
  user_tags: "Mention users",
  subreddit: "Target subreddit",
  post_type: "Post type",
  nsfw: "NSFW flag",
  board_id: "Board to pin",
  source_type: "Media source type",
  author: "Actor URN",
  lifecycle_state: "Publication state",
  visibility: "Post visibility"
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
