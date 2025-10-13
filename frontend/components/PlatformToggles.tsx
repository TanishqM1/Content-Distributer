"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Switch } from "@/components/ui/switch";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { Platform, PLATFORM_NAMES } from "@/lib/types";
import { cn } from "@/lib/utils";
import { 
  Youtube, 
  MessageSquare, 
  Instagram, 
  Pin, 
  Linkedin 
} from "lucide-react";

interface PlatformTogglesProps {
  selectedPlatforms: Platform[];
  onPlatformToggle: (platform: Platform, checked: boolean) => void;
}

// Map each platform to its corresponding icon from Lucide React
const platformIcons = {
  youtube: Youtube,
  reddit: MessageSquare,
  instagram: Instagram,
  pinterest: Pin,
  linkedin: Linkedin,
};

/**
 * Component that lets users choose which social media platforms they want to post to
 * Each platform is shown as a card with an icon, name, description, and toggle switch
 */
export function PlatformToggles({ selectedPlatforms, onPlatformToggle }: PlatformTogglesProps) {
  // All the platforms we support
  const platforms: Platform[] = ["youtube", "reddit", "instagram", "pinterest", "linkedin"];

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span>Distribution Platforms</span>
          {/* Show how many platforms the user has selected */}
          <Badge variant="secondary">
            {selectedPlatforms.length} selected
          </Badge>
        </CardTitle>
      </CardHeader>
      <CardContent>
        {/* Display all platforms in a responsive grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {platforms.map((platform) => {
            const Icon = platformIcons[platform];
            const isSelected = selectedPlatforms.includes(platform);
            
            return (
              <div
                key={platform}
                className={cn(
                  "relative p-4 border rounded-lg transition-all duration-200 cursor-pointer hover:shadow-md",
                  // Different styling for selected vs unselected platforms
                  isSelected
                    ? "border-primary bg-primary/5 shadow-sm"
                    : "border-muted hover:border-muted-foreground/50"
                )}
                onClick={() => onPlatformToggle(platform, !isSelected)}
              >
                <div className="flex items-center space-x-3">
                  {/* Platform icon with different colors based on selection state */}
                  <div className={cn(
                    "w-10 h-10 rounded-lg flex items-center justify-center",
                    isSelected ? "bg-primary text-primary-foreground" : "bg-muted"
                  )}>
                    <Icon className="w-5 h-5" />
                  </div>
                  <div className="flex-1 min-w-0">
                    {/* Platform name */}
                    <Label 
                      htmlFor={platform} 
                      className="text-sm font-medium cursor-pointer"
                    >
                      {PLATFORM_NAMES[platform]}
                    </Label>
                    {/* Short description of what this platform is for */}
                    <p className="text-xs text-muted-foreground mt-1">
                      {platform === "youtube" && "Video content & live streams"}
                      {platform === "reddit" && "Community discussions & posts"}
                      {platform === "instagram" && "Visual content & stories"}
                      {platform === "pinterest" && "Visual discovery & inspiration"}
                      {platform === "linkedin" && "Professional networking"}
                    </p>
                  </div>
                  {/* Toggle switch - clicking it or the card will toggle the platform */}
                  <Switch
                    id={platform}
                    checked={isSelected}
                    onCheckedChange={(checked) => onPlatformToggle(platform, checked)}
                    onClick={(e) => e.stopPropagation()} // Prevent double-triggering when clicking the switch
                  />
                </div>
              </div>
            );
          })}
        </div>
        
        {/* Show a message if no platforms are selected */}
        {selectedPlatforms.length === 0 && (
          <div className="text-center py-8 text-muted-foreground">
            <p className="text-sm">Select at least one platform to continue</p>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
