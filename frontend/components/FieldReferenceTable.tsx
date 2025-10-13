"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { FIELD_DESCRIPTIONS, FIELD_APPLIES_TO, PLATFORM_NAMES } from "@/lib/types";

export function FieldReferenceTable() {
  const fieldData = [
    {
      field: "title",
      purpose: "Title of post/video",
      appliesTo: ["youtube", "reddit", "pinterest"],
      notes: "Required on all three"
    },
    {
      field: "description", 
      purpose: "Long description/body",
      appliesTo: ["youtube", "pinterest", "linkedin"],
      notes: "Optional for LinkedIn (as text)"
    },
    {
      field: "caption",
      purpose: "Short caption / post text", 
      appliesTo: ["instagram", "linkedin"],
      notes: "Hashtags allowed"
    },
    {
      field: "tags",
      purpose: "Keywords array",
      appliesTo: ["youtube"],
      notes: "Optional"
    },
    {
      field: "image_url",
      purpose: "Image URL",
      appliesTo: ["instagram", "pinterest", "reddit", "linkedin"],
      notes: "Required if image post"
    },
    {
      field: "thumbnail",
      purpose: "Thumbnail image",
      appliesTo: [],
      notes: "Not used (TikTok removed)"
    },
    {
      field: "video_file",
      purpose: "Video file URL/path",
      appliesTo: ["youtube", "linkedin", "instagram"],
      notes: "Required for YouTube uploads"
    },
    {
      field: "link",
      purpose: "External URL",
      appliesTo: ["reddit", "pinterest"],
      notes: "Optional otherwise"
    },
    {
      field: "content",
      purpose: "Text body (self post)",
      appliesTo: ["reddit"],
      notes: "Required if post_type=self"
    },
    {
      field: "category_id",
      purpose: "YouTube category",
      appliesTo: ["youtube"],
      notes: "Optional"
    },
    {
      field: "privacy_status",
      purpose: "Visibility",
      appliesTo: ["youtube"],
      notes: "public/unlisted/private"
    },
    {
      field: "location_id",
      purpose: "Location tag",
      appliesTo: ["instagram"],
      notes: "Optional"
    },
    {
      field: "user_tags",
      purpose: "Mention users",
      appliesTo: ["instagram"],
      notes: "Optional"
    },
    {
      field: "subreddit",
      purpose: "Target subreddit",
      appliesTo: ["reddit"],
      notes: "Required"
    },
    {
      field: "post_type",
      purpose: "Post type",
      appliesTo: ["reddit"],
      notes: "self/link/image"
    },
    {
      field: "nsfw",
      purpose: "NSFW flag",
      appliesTo: ["reddit"],
      notes: "Optional"
    },
    {
      field: "board_id",
      purpose: "Board to pin",
      appliesTo: ["pinterest"],
      notes: "Required"
    },
    {
      field: "source_type",
      purpose: "Media source type",
      appliesTo: ["pinterest"],
      notes: "Usually image_url"
    },
    {
      field: "author",
      purpose: "Actor URN",
      appliesTo: ["linkedin"],
      notes: "Required"
    },
    {
      field: "lifecycle_state",
      purpose: "Publication state",
      appliesTo: ["linkedin"],
      notes: "Default PUBLISHED"
    },
    {
      field: "visibility",
      purpose: "Post visibility",
      appliesTo: ["linkedin"],
      notes: "PUBLIC or restricted"
    }
  ];

  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg">Field Reference</CardTitle>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[120px]">Field</TableHead>
              <TableHead className="w-[200px]">Purpose</TableHead>
              <TableHead className="w-[200px]">Applies To</TableHead>
              <TableHead>Notes</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {fieldData.map((row) => (
              <TableRow key={row.field}>
                <TableCell className="font-medium">
                  <code className="text-xs bg-muted px-1 py-0.5 rounded">
                    {row.field}
                  </code>
                </TableCell>
                <TableCell className="text-sm">{row.purpose}</TableCell>
                <TableCell>
                  <div className="flex flex-wrap gap-1">
                    {row.appliesTo.map((platform) => (
                      <Badge key={platform} variant="outline" className="text-xs">
                        {PLATFORM_NAMES[platform]}
                      </Badge>
                    ))}
                  </div>
                </TableCell>
                <TableCell className="text-sm text-muted-foreground">
                  {row.notes}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
}
