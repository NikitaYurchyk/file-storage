import { useEffect, useState } from "react";
import {
  deleteFile,
  downloadFile,
  fetchFile,
  fetchFiles,
  uploadFile,
} from "./api/files";
import "./App.css";
import { FileCard } from "./components/FileCard";
import { VideoPlayer } from "./components/VideoPlayer";
import type { FileItem, FileResponse } from "./types/file";

function App() {
  const [files, setFiles] = useState<FileItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [currentVideo, setCurrentVideo] = useState<FileResponse | null>(null);
  const [uploading, setUploading] = useState(false);
  const [searchQuery, setSearchQuery] = useState("");
  const viewMode = "grid";

  const filteredFiles = files.filter((file) => {
    const query = searchQuery.toLowerCase();
    const filename = (file.filename || "").toLowerCase();
    const title = (file.title || "").toLowerCase();
    return filename.includes(query) || title.includes(query);
  });

  const loadFiles = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await fetchFiles();
      setFiles(data || []);
    } catch (err) {
      setError("Failed to load files");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadFiles();
  }, []);

  const handlePlay = async (file: FileItem) => {
    try {
      const fileData = await fetchFile(file.id);
      setCurrentVideo(fileData);
    } catch (err) {
      console.error("Failed to get video URL:", err);
    }
  };

  const handleDelete = async (id: string) => {
    if (!confirm("Move to trash?")) return;
    try {
      await deleteFile(id);
      setFiles(files.filter((f) => f.id !== id));
    } catch (err) {
      console.error("Failed to delete file:", err);
    }
  };

  const handleDownload = async (file: FileItem) => {
    try {
      await downloadFile(file.id);
    } catch (err) {
      console.error("Failed to download file:", err);
    }
  };

  const handleUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    try {
      setUploading(true);
      await uploadFile(file);
      await loadFiles();
    } catch (err) {
      console.error("Failed to upload file:", err);
    } finally {
      setUploading(false);
      e.target.value = "";
    }
  };

  return (
    <div className="app">
      <header className="header">
        <div className="header-left">
          <span className="logo-text">OpenDrive</span>
        </div>
        <></>
        <div className="search-bar">
          <svg
            className="search-icon"
            viewBox="0 0 24 24"
            width="20"
            height="20"
          >
            <path
              fill="#5f6368"
              d="M15.5 14h-.79l-.28-.27A6.471 6.471 0 0016 9.5 6.5 6.5 0 109.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"
            />
          </svg>
          <input
            type="text"
            placeholder="Search in Drive"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
        </div>
        <div className="header-right"></div>
      </header>

      <div className="layout">
        <aside className="sidebar">
          <label className="new-btn">
            <svg viewBox="0 0 24 24" width="20" height="20">
              <path fill="#5f6368" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z" />
            </svg>
            <span>{uploading ? "Uploading..." : "New"}</span>
            <input
              type="file"
              onChange={handleUpload}
              disabled={uploading}
              hidden
            />
          </label>

          <nav className="nav">
            <a className="nav-item active">
              {/* <svg viewBox="0 0 24 24" width="20" height="20">
                <path
                  fill="currentColor"
                  d="M19 2H5a2 2 0 00-2 2v16a2 2 0 002 2h14a2 2 0 002-2V4a2 2 0 00-2-2zm0 18H5V4h14v16z"
                />
              </svg> */}
              My Drive
            </a>
            
          </nav>

          <div className="storage-info">
            <div className="storage-bar">
              <div className="storage-used" style={{ width: "23%" }}></div>
            </div>
            <span className="storage-text">2.3 GB of 15 GB used</span>
          </div>
        </aside>

        <main className="main">
          {loading && <div className="loading">Loading...</div>}
          {error && <div className="error">{error}</div>}
          {!loading && !error && files.length === 0 && (
            <div className="empty">
              <svg viewBox="0 0 24 24" width="120" height="120">
                <path
                  fill="#dadce0"
                  d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14z"
                />
              </svg>
              <h3>Drop files here</h3>
              <p>or use the "New" button</p>
            </div>
          )}
          {!loading &&
            !error &&
            files.length > 0 &&
            filteredFiles.length === 0 && (
              <div className="empty">
                <h3>No results found</h3>
                <p>Try a different search term</p>
              </div>
            )}
          <div className={`file-grid ${viewMode}`}>
            {filteredFiles.map((file) => (
              <FileCard
                key={file.id}
                file={file}
                viewMode={viewMode}
                onPlay={handlePlay}
                onDelete={handleDelete}
                onDownload={handleDownload}
              />
            ))}
          </div>
        </main>
      </div>

      {currentVideo && (
        <VideoPlayer
          url={currentVideo.signed_url}
          title={currentVideo.title || currentVideo.filename}
          onClose={() => setCurrentVideo(null)}
        />
      )}
    </div>
  );
}

export default App;
