export interface FileItem {
  id: string;
  filename: string;
  file_url: string;
  file_type: string;
  size: number;
  aspect_ratio?: string;
  title?: string;
  description?: string;
  thumbnail_url?: string;
  created_at: number;
  updated_at: number;
}

export interface FileResponse {
  id: string;
  filename: string;
  content_type: string;
  size: number;
  signed_url: string;
  aspect_ratio?: string;
  title?: string;
  description?: string;
  created_at: number;
  expires_at: string;
}
