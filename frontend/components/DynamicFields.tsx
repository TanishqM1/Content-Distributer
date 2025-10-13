"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Badge } from "@/components/ui/badge";
import { Label } from "@/components/ui/label";
import { Platform, FieldKey, FIELD_DESCRIPTIONS, FIELD_APPLIES_TO, FormData } from "@/lib/types";
import { getPlatformsForField } from "@/lib/helpers";
import { UseFormRegister, FieldErrors, UseFormWatch } from "react-hook-form";

interface DynamicFieldsProps {
  selectedPlatforms: Platform[];
  requiredFields: FieldKey[];
  register: UseFormRegister<FormData>;
  errors: FieldErrors<FormData>;
  watch: UseFormWatch<FormData>;
}

/**
 * Component that dynamically creates form fields based on which platforms the user selected
 * Each field shows which platforms need it and has appropriate input types (text, dropdown, etc.)
 */
export function DynamicFields({ 
  selectedPlatforms, 
  requiredFields, 
  register, 
  errors, 
  watch 
}: DynamicFieldsProps) {
  const watchedValues = watch();

  /**
   * Creates the right type of input field for each field type
   * Some fields are text inputs, others are dropdowns, some are textareas
   */
  const renderField = (field: FieldKey) => {
    const platforms = getPlatformsForField(field, selectedPlatforms);
    const error = errors[field];
    const value = watchedValues[field];

    // Set up form registration and error styling
    const fieldProps = {
      ...register(field),
      className: error ? "border-red-500" : "",
    };

    // Create label with platform badges showing which platforms need this field
    const label = (
      <div className="flex items-center gap-2">
        <Label htmlFor={field} className="text-sm font-medium capitalize">
          {field.replace(/_/g, " ")}
        </Label>
        <div className="flex gap-1">
          {platforms.map((platform) => (
            <Badge key={platform} variant="secondary" className="text-xs">
              {platform}
            </Badge>
          ))}
        </div>
      </div>
    );

    // Helper text explaining what this field is for
    const helperText = (
      <div className="text-xs text-muted-foreground">
        {FIELD_DESCRIPTIONS[field]}
      </div>
    );

    switch (field) {
      // Simple text fields - just need a regular input
      case "title":
      case "description":
      case "caption":
      case "subreddit":
      case "board_id":
      case "author":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Input
              {...fieldProps}
              id={field}
              placeholder={`Enter ${field.replace(/_/g, " ")}`}
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Multi-line text field for longer content (like Reddit posts)
      case "content":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Textarea
              {...fieldProps}
              id={field}
              placeholder="Enter post content"
              rows={4}
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // URL fields - need to be valid URLs
      case "image_url":
      case "thumbnail":
      case "video_file":
      case "link":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Input
              {...fieldProps}
              id={field}
              type="url"
              placeholder={`Enter ${field.replace(/_/g, " ")} URL`}
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Tags field - user enters comma-separated values
      case "tags":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Input
              {...fieldProps}
              id={field}
              placeholder="Enter tags separated by commas"
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // YouTube privacy settings dropdown
      case "privacy_status":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Select {...fieldProps}>
              <SelectTrigger>
                <SelectValue placeholder="Select privacy status" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="public">Public</SelectItem>
                <SelectItem value="unlisted">Unlisted</SelectItem>
                <SelectItem value="private">Private</SelectItem>
              </SelectContent>
            </Select>
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Reddit post type dropdown
      case "post_type":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Select {...fieldProps}>
              <SelectTrigger>
                <SelectValue placeholder="Select post type" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="self">Self Post</SelectItem>
                <SelectItem value="link">Link Post</SelectItem>
                <SelectItem value="image">Image Post</SelectItem>
              </SelectContent>
            </Select>
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // LinkedIn visibility dropdown
      case "visibility":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Select {...fieldProps}>
              <SelectTrigger>
                <SelectValue placeholder="Select visibility" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="PUBLIC">Public</SelectItem>
                <SelectItem value="CONNECTIONS">Connections Only</SelectItem>
              </SelectContent>
            </Select>
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Reddit NSFW flag dropdown
      case "nsfw":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Select {...fieldProps}>
              <SelectTrigger>
                <SelectValue placeholder="Select NSFW status" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="false">No</SelectItem>
                <SelectItem value="true">Yes</SelectItem>
              </SelectContent>
            </Select>
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Optional text fields for various platform-specific settings
      case "category_id":
      case "location_id":
      case "source_type":
      case "lifecycle_state":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Input
              {...fieldProps}
              id={field}
              placeholder={`Enter ${field.replace(/_/g, " ")}`}
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Instagram user tags - mention other users (comma-separated)
      case "user_tags":
        return (
          <div key={field} className="space-y-2">
            {label}
            <Input
              {...fieldProps}
              id={field}
              placeholder="Enter user tags separated by commas"
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );

      // Fallback for any field we didn't specifically handle
      default:
        return (
          <div key={field} className="space-y-2">
            {label}
            <Input
              {...fieldProps}
              id={field}
              placeholder={`Enter ${String(field).replace(/_/g, " ")}`}
            />
            {helperText}
            {error && (
              <p className="text-sm text-red-500">{error.message}</p>
            )}
          </div>
        );
    }
  };

  // Check if Instagram is selected and if user has provided media
  const showInstagramMediaFields = selectedPlatforms.includes("instagram");
  const hasImageUrl = watchedValues.image_url;
  const hasVideoFile = watchedValues.video_file;

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span>Content Configuration</span>
          {/* Show how many fields the user needs to fill out */}
          <Badge variant="outline">
            {requiredFields.length} fields
          </Badge>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-6">
        {requiredFields.length > 0 ? (
          // Display all required fields in a responsive grid
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {requiredFields.map(renderField)}
          </div>
        ) : (
          // Show message if no additional fields are needed
          <div className="text-center py-8 text-muted-foreground">
            <p className="text-sm">No additional fields required for selected platforms</p>
          </div>
        )}
        
        {/* Special warning for Instagram - it requires either an image or video */}
        {showInstagramMediaFields && (
          <div className="p-4 border rounded-lg bg-blue-50 dark:bg-blue-950/20 border-blue-200 dark:border-blue-800">
            <div className="flex items-start space-x-3">
              <div className="w-5 h-5 rounded-full bg-blue-100 dark:bg-blue-900 flex items-center justify-center flex-shrink-0 mt-0.5">
                <span className="text-blue-600 dark:text-blue-400 text-xs font-bold">i</span>
              </div>
              <div>
                <p className="text-sm font-medium text-blue-900 dark:text-blue-100 mb-1">
                  Instagram Media Requirement
                </p>
                <p className="text-xs text-blue-700 dark:text-blue-300">
                  Instagram requires either an image or video. 
                  {!hasImageUrl && !hasVideoFile && (
                    <span className="text-red-500 ml-1 font-medium">Please provide at least one media field.</span>
                  )}
                </p>
              </div>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
