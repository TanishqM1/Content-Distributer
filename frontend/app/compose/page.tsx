"use client";

import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Platform, FormData } from "@/lib/types";
import { getRequiredFields, buildZodSchema } from "@/lib/helpers";
import { PlatformToggles } from "@/components/PlatformToggles";
import { DynamicFields } from "@/components/DynamicFields";
import { SummaryPanel } from "@/components/SummaryPanel";
import { FieldReferenceTable } from "@/components/FieldReferenceTable";
import { VideoUpload } from "@/components/VideoUpload";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { useToast } from "@/components/ui/use-toast";
import { Toaster } from "@/components/ui/toaster";
import { ArrowLeft, Settings, Save, FileVideo } from "lucide-react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";

export default function ComposePage() {
  const [selectedPlatforms, setSelectedPlatforms] = useState<Platform[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [selectedVideo, setSelectedVideo] = useState<File | null>(null);
  const [currentStep, setCurrentStep] = useState<"platforms" | "content" | "review">("platforms");
  const { toast } = useToast();

  const requiredFields = getRequiredFields(selectedPlatforms);
  const schema = buildZodSchema(requiredFields, selectedPlatforms);

  const {
    register,
    handleSubmit,
    watch,
    setValue,
    formState: { errors },
    reset
  } = useForm<FormData>({
    resolver: zodResolver(schema),
    mode: "onChange"
  });

  const formData = watch();

  const handlePlatformToggle = (platform: Platform, checked: boolean) => {
    if (checked) {
      setSelectedPlatforms(prev => [...prev, platform]);
    } else {
      setSelectedPlatforms(prev => prev.filter(p => p !== platform));
      reset();
    }
  };

  const handleVideoSelect = async (file: File | null) => {
    setSelectedVideo(file);
    if (file) {
      // Upload file to backend
      try {
        const formData = new globalThis.FormData();
        formData.append('file', file);
        
        const response = await fetch('http://localhost:8000/upload/file', {
          method: 'POST',
          body: formData,
        });
        
        if (response.ok) {
          const result = await response.json();
          console.log('File uploaded successfully:', result);
          // Store the file path and filename for later use
          const fileWithMetadata = Object.assign(file, {
            path: result.file_path,
            savedName: result.filename
          });
          setSelectedVideo(fileWithMetadata);
          
          // Auto-populate image_url with the uploaded file path
          setValue('image_url', result.file_path);
          
          toast({
            title: "File Uploaded!",
            description: `"${file.name}" was uploaded successfully.`,
          });
        } else {
          console.error('File upload failed');
          toast({
            title: "File Upload Failed",
            description: "Failed to upload file. Please try again.",
            variant: "destructive",
          });
        }
      } catch (error) {
        console.error('Error uploading file:', error);
        toast({
          title: "Error Uploading File",
          description: "An unexpected error occurred during file upload. Please try again.",
          variant: "destructive",
        });
      }
    }
  };

  const onSubmit = async (data: any) => {
    setIsSubmitting(true);
    try {
      console.log("Submitting data:", data);
      
      // Prepare the payload for the backend
      const payload = {
        platforms: selectedPlatforms,
        ...data,
        // user_tags is already a string from the form input
        user_tags: data.user_tags || '',
        // Add video file path if a file was selected
        media_file: selectedVideo && (selectedVideo as any).savedName ? (selectedVideo as any).savedName : ''
      };
      
      // Send to backend
      const response = await fetch('http://localhost:8000/post/content', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });
      
      if (response.ok) {
        const result = await response.json();
        console.log("Backend response:", result);
        toast({
          title: "Content Submitted!",
          description: `Successfully posted to ${result.platforms?.join(', ') || 'selected platforms'}.`,
        });
      } else {
        const error = await response.text();
        console.error("Backend error:", error);
        toast({
          title: "Error Submitting Content",
          description: `Backend error: ${error}`,
          variant: "destructive",
        });
      }
    } catch (error) {
      console.error("Error submitting content:", error);
      toast({
        title: "Submission Failed",
        description: "An unexpected error occurred. Please try again.",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  const steps = [
    { id: "platforms", label: "Platforms", description: "Select distribution channels" },
    { id: "content", label: "Content", description: "Upload and configure content" },
    { id: "review", label: "Review", description: "Review and publish" }
  ];

  return (
    <div className="min-h-screen bg-muted/30">
      {/* Header */}
      <div className="border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <Link href="/">
                <Button variant="ghost" size="sm">
                  <ArrowLeft className="w-4 h-4 mr-2" />
                  Back
                </Button>
              </Link>
              <div>
                <h1 className="text-xl font-semibold">Content Composer</h1>
                <p className="text-sm text-muted-foreground">
                  Create and distribute content across platforms
                </p>
              </div>
            </div>
            <div className="flex items-center space-x-2">
              <Button variant="outline" size="sm">
                <Save className="w-4 h-4 mr-2" />
                Save Draft
              </Button>
              <Button variant="outline" size="sm">
                <Settings className="w-4 h-4" />
              </Button>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-4 py-6">
        {/* Progress Steps */}
        <div className="mb-8">
          <div className="flex items-center justify-between">
            {steps.map((step, index) => (
              <div key={step.id} className="flex items-center">
                <div className="flex flex-col items-center">
                  <div className={cn(
                    "w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium",
                    currentStep === step.id 
                      ? "bg-primary text-primary-foreground" 
                      : "bg-muted text-muted-foreground"
                  )}>
                    {index + 1}
                  </div>
                  <div className="mt-2 text-center">
                    <p className="text-sm font-medium">{step.label}</p>
                    <p className="text-xs text-muted-foreground">{step.description}</p>
                  </div>
                </div>
                {index < steps.length - 1 && (
                  <div className="w-16 h-px bg-muted mx-4" />
                )}
              </div>
            ))}
          </div>
        </div>

        <div className="grid grid-cols-1 xl:grid-cols-4 gap-6">
          {/* Main Content Area */}
          <div className="xl:col-span-3 space-y-6">
            {/* Step 1: Platform Selection */}
            {currentStep === "platforms" && (
              <PlatformToggles
                selectedPlatforms={selectedPlatforms}
                onPlatformToggle={handlePlatformToggle}
              />
            )}

            {/* Step 2: Content Creation */}
            {currentStep === "content" && (
              <div className="space-y-6">
                <VideoUpload
                  onVideoSelect={handleVideoSelect}
                  selectedVideo={selectedVideo}
                />

                {selectedPlatforms.length > 0 && (
                  <DynamicFields
                    selectedPlatforms={selectedPlatforms}
                    requiredFields={requiredFields}
                    register={register}
                    errors={errors}
                    watch={watch}
                  />
                )}
              </div>
            )}

            {/* Step 3: Review */}
            {currentStep === "review" && (
              <div className="space-y-6">
                <Card>
                  <CardHeader>
                    <CardTitle>Content Preview</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-4">
                      {selectedVideo && (
                        <div className="flex items-center space-x-3 p-3 border rounded-lg">
                          <div className="w-12 h-12 bg-primary/10 rounded flex items-center justify-center">
                            <FileVideo className="w-6 h-6 text-primary" />
                          </div>
                          <div>
                            <p className="font-medium">{selectedVideo.name}</p>
                            <p className="text-sm text-muted-foreground">
                              {(selectedVideo.size / (1024 * 1024)).toFixed(1)} MB
                            </p>
                          </div>
                        </div>
                      )}
                      
                      <div className="space-y-2">
                        <h4 className="font-medium">Selected Platforms</h4>
                        <div className="flex flex-wrap gap-2">
                          {selectedPlatforms.map((platform) => (
                            <Badge key={platform} variant="secondary">
                              {platform}
                            </Badge>
                          ))}
                        </div>
                      </div>
                    </div>
                  </CardContent>
                </Card>

                <FieldReferenceTable />
              </div>
            )}

            {/* Navigation Buttons */}
            <div className="flex justify-between">
              <Button
                variant="outline"
                onClick={() => {
                  if (currentStep === "content") setCurrentStep("platforms");
                  if (currentStep === "review") setCurrentStep("content");
                }}
                disabled={currentStep === "platforms"}
              >
                Previous
              </Button>
              
              <Button
                onClick={() => {
                  if (currentStep === "platforms") setCurrentStep("content");
                  if (currentStep === "content") setCurrentStep("review");
                  if (currentStep === "review") handleSubmit(onSubmit)();
                }}
                disabled={currentStep === "review" || selectedPlatforms.length === 0}
              >
                {currentStep === "review" ? "Publish" : "Next"}
              </Button>
            </div>
          </div>

          {/* Sidebar */}
          <div className="xl:col-span-1">
            <SummaryPanel
              selectedPlatforms={selectedPlatforms}
              formData={formData}
              errors={errors}
              onSubmit={onSubmit}
              isSubmitting={isSubmitting}
            />
          </div>
        </div>
      </div>
      <Toaster />
    </div>
  );
}
