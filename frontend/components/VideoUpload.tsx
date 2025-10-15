"use client";

import { useState, useCallback } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Upload, X, FileVideo, CheckCircle, Image } from "lucide-react";
import { cn } from "@/lib/utils";

interface VideoUploadProps {
  onVideoSelect: (file: File | null) => void;
  selectedVideo?: File | null;
  className?: string;
}

export function VideoUpload({ onVideoSelect, selectedVideo, className }: VideoUploadProps) {
  const [isDragOver, setIsDragOver] = useState(false);

  const handleDragOver = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDragOver(true);
  }, []);

  const handleDragLeave = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDragOver(false);
  }, []);

  const handleDrop = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDragOver(false);
    
    const files = Array.from(e.dataTransfer.files);
    const mediaFile = files.find(file => 
      file.type.startsWith('video/') || file.type.startsWith('image/')
    );
    
    if (mediaFile) {
      onVideoSelect(mediaFile);
    }
  }, [onVideoSelect]);

  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file && (file.type.startsWith('video/') || file.type.startsWith('image/'))) {
      onVideoSelect(file);
    }
  };

  const handleRemove = () => {
    onVideoSelect(null);
  };

  return (
    <Card className={cn("w-full", className)}>
      <CardContent className="p-6">
        <div className="space-y-4">
          <div className="text-center">
            <h3 className="text-lg font-semibold mb-2">Media Upload</h3>
            <p className="text-sm text-muted-foreground">
              Upload your video or image file for distribution across platforms
            </p>
          </div>

          {!selectedVideo ? (
            <div
              className={cn(
                "border-2 border-dashed rounded-lg p-8 text-center transition-colors",
                isDragOver
                  ? "border-primary bg-primary/5"
                  : "border-muted-foreground/25 hover:border-muted-foreground/50"
              )}
              onDragOver={handleDragOver}
              onDragLeave={handleDragLeave}
              onDrop={handleDrop}
            >
              <div className="flex flex-col items-center space-y-4">
                <div className="w-16 h-16 rounded-full bg-muted flex items-center justify-center">
                  <Upload className="w-8 h-8 text-muted-foreground" />
                </div>
                <div className="space-y-2">
                  <p className="text-lg font-medium">
                    Drag and drop media files to upload
                  </p>
                  <p className="text-sm text-muted-foreground">
                    Your media will be processed for distribution
                  </p>
                </div>
                <div className="flex flex-col sm:flex-row gap-3">
                  <Button variant="outline" asChild>
                    <label htmlFor="media-upload" className="cursor-pointer">
                      <FileVideo className="w-4 h-4 mr-2" />
                      Select video file
                    </label>
                  </Button>
                  <Button variant="outline" asChild>
                    <label htmlFor="media-upload" className="cursor-pointer">
                      <Image className="w-4 h-4 mr-2" />
                      Select image file
                    </label>
                  </Button>
                  <input
                    id="media-upload"
                    type="file"
                    accept="video/*,image/*"
                    onChange={handleFileSelect}
                    className="hidden"
                  />
                </div>
              </div>
            </div>
          ) : (
            <div className="border rounded-lg p-4 bg-muted/50">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 rounded bg-primary/10 flex items-center justify-center">
                    {selectedVideo?.type?.startsWith('video/') ? (
                      <FileVideo className="w-5 h-5 text-primary" />
                    ) : (
                      <Image className="w-5 h-5 text-primary" />
                    )}
                  </div>
                  <div className="flex-1 min-w-0">
                    <p className="text-sm font-medium truncate">
                      {selectedVideo?.name || 'Unknown file'}
                    </p>
                    <p className="text-xs text-muted-foreground">
                      {selectedVideo?.size ? `${(selectedVideo.size / (1024 * 1024)).toFixed(1)} MB` : 'Unknown size'}
                    </p>
                  </div>
                </div>
                <div className="flex items-center space-x-2">
                  <CheckCircle className="w-4 h-4 text-green-500" />
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={handleRemove}
                    className="h-8 w-8 p-0"
                  >
                    <X className="w-4 h-4" />
                  </Button>
                </div>
              </div>
            </div>
          )}

          <div className="text-xs text-muted-foreground text-center">
            Supported formats: MP4, MOV, AVI, WMV, FLV, WebM, JPG, PNG, GIF
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
