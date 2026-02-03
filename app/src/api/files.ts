import type { FileItem, FileResponse } from '../types/file';

const API_BASE = '/api/v1';

export async function fetchFiles(): Promise<FileItem[]> {
  const response = await fetch(`${API_BASE}/files`);
  if (!response.ok) {
    throw new Error('Failed to fetch files');
  }
  return response.json();
}

export async function fetchFile(id: string): Promise<FileResponse> {
  const response = await fetch(`${API_BASE}/files/${id}`);
  if (!response.ok) {
    throw new Error('Failed to fetch file');
  }
  return response.json();
}

export async function deleteFile(id: string): Promise<void> {
  const response = await fetch(`${API_BASE}/files/${id}`, {
    method: 'DELETE',
  });
  if (!response.ok) {
    throw new Error('Failed to delete file');
  }
}

export async function downloadFile(id: string): Promise<void> {
  const fileData = await fetchFile(id);
  window.open(fileData.signed_url, '_blank');
}

export async function uploadFile(file: File, title?: string, description?: string): Promise<void> {
  const formData = new FormData();
  formData.append('file', file);
  if (title) formData.append('title', title);
  if (description) formData.append('description', description);

  const response = await fetch(`${API_BASE}/files`, {
    method: 'POST',
    body: formData,
  });
  if (!response.ok) {
    throw new Error('Failed to upload file');
  }
}
