"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Platform, PLATFORM_NAMES, FormData } from "@/lib/types";
import { transformToSubmitData } from "@/lib/helpers";
import { FieldErrors } from "react-hook-form";
import { useState } from "react";

interface SummaryPanelProps {
  selectedPlatforms: Platform[];
  formData: FormData;
  errors: FieldErrors<FormData>;
  onSubmit: (data: any) => void;
  isSubmitting: boolean;
}

export function SummaryPanel({ 
  selectedPlatforms, 
  formData, 
  errors, 
  onSubmit, 
  isSubmitting 
}: SummaryPanelProps) {
  const [showJson, setShowJson] = useState(false);
  
  const submitData = transformToSubmitData(formData, selectedPlatforms);
  const errorFields = Object.keys(errors);
  const hasErrors = errorFields.length > 0;

  const handleSubmit = () => {
    onSubmit(submitData);
  };

  return (
    <Card className="h-fit sticky top-6">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span>Publish Summary</span>
          <Badge variant={hasErrors ? "destructive" : "secondary"}>
            {hasErrors ? "Issues" : "Ready"}
          </Badge>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* Selected Platforms */}
        <div>
          <h4 className="text-sm font-medium mb-2">Selected Platforms</h4>
          <div className="flex flex-wrap gap-2">
            {selectedPlatforms.length > 0 ? (
              selectedPlatforms.map((platform) => (
                <Badge key={platform} variant="default">
                  {PLATFORM_NAMES[platform]}
                </Badge>
              ))
            ) : (
              <p className="text-sm text-muted-foreground">No platforms selected</p>
            )}
          </div>
        </div>

        {/* Validation Errors */}
        {hasErrors && (
          <div>
            <h4 className="text-sm font-medium mb-2 text-red-600">Missing Required Fields</h4>
            <div className="space-y-1">
              {errorFields.map((field) => (
                <div key={field} className="text-sm text-red-600">
                  • {String(field).replace(/_/g, " ")}: {errors[field]?.message}
                </div>
              ))}
            </div>
          </div>
        )}

        {/* JSON Preview Toggle */}
        <div>
          <Button
            variant="outline"
            size="sm"
            onClick={() => setShowJson(!showJson)}
            className="w-full"
          >
            {showJson ? "Hide" : "Show"} JSON Preview
          </Button>
        </div>

        {/* JSON Preview */}
        {showJson && (
          <div>
            <h4 className="text-sm font-medium mb-2">JSON Payload</h4>
            <pre className="bg-muted p-3 rounded-md text-xs overflow-auto max-h-64">
              {JSON.stringify(submitData, null, 2)}
            </pre>
          </div>
        )}

        {/* Submit Button */}
        <Button
          onClick={handleSubmit}
          disabled={hasErrors || selectedPlatforms.length === 0 || isSubmitting}
          className="w-full"
        >
          {isSubmitting ? "Submitting..." : "Submit Content"}
        </Button>

        {selectedPlatforms.length === 0 && (
          <p className="text-sm text-muted-foreground text-center">
            Select at least one platform to continue
          </p>
        )}
      </CardContent>
    </Card>
  );
}
