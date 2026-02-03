import type { FileItem } from '../types/file';

interface FileCardProps {
  file: FileItem;
  viewMode: 'grid' | 'list';
  onPlay: (file: FileItem) => void;
  onDelete: (id: string) => void;
  onDownload: (file: FileItem) => void;
}

export function FileCard({ file, viewMode, onPlay, onDelete, onDownload }: FileCardProps) {
  const isVideo = file.file_type?.startsWith('video/');
  const isImage = file.file_type?.startsWith('image/');
  const isPdf = file.file_type === 'application/pdf';

  const formatSize = (bytes: number) => {
    if (!bytes || bytes === 0) return '-';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
  };

  const formatDate = (timestamp: number) => {
    if (!timestamp) return '-';
    const date = new Date(timestamp * 1000);
    const now = new Date();
    const diff = now.getTime() - date.getTime();
    const days = Math.floor(diff / (1000 * 60 * 60 * 24));

    if (days === 0) return 'Today';
    if (days === 1) return 'Yesterday';
    if (days < 7) return `${days} days ago`;
    return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });
  };

  const getFileIcon = () => {
    if (isVideo) {
      return (
        <svg viewBox="0 0 24 24" width="24" height="24">
          <path fill="#ea4335" d="M18 4l2 4h-3l-2-4h-2l2 4h-3l-2-4H8l2 4H7L5 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4h-4z"/>
        </svg>
      );
    }
    if (isImage) {
      return (
        <svg viewBox="0 0 24 24" width="24" height="24">
          <path fill="#ea4335" d="M21 19V5c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zM8.5 13.5l2.5 3.01L14.5 12l4.5 6H5l3.5-4.5z"/>
        </svg>
      );
    }
    if (isPdf) {
      return (
        <svg viewBox="0 0 24 24" width="24" height="24">
          <path fill="#ea4335" d="M20 2H8c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-8.5 7.5c0 .83-.67 1.5-1.5 1.5H9v2H7.5V7H10c.83 0 1.5.67 1.5 1.5v1zm5 2c0 .83-.67 1.5-1.5 1.5h-2.5V7H15c.83 0 1.5.67 1.5 1.5v3zm4-3H19v1h1.5V11H19v2h-1.5V7h3v1.5zM9 9.5h1v-1H9v1zM4 6H2v14c0 1.1.9 2 2 2h14v-2H4V6zm10 5.5h1v-3h-1v3z"/>
        </svg>
      );
    }
    return (
      <svg viewBox="0 0 24 24" width="24" height="24">
        <path fill="#5f6368" d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zm2 16H8v-2h8v2zm0-4H8v-2h8v2zm-3-5V3.5L18.5 9H13z"/>
      </svg>
    );
  };

  const getThumbnailColor = () => {
    if (isVideo) return '#fce8e6';
    if (isImage) return '#e8f0fe';
    if (isPdf) return '#fce8e6';
    return '#f1f3f4';
  };

  const handleClick = () => {
    if (isVideo) {
      onPlay(file);
    }
  };

  if (viewMode === 'list') {
    return (
      <div className="file-row" onClick={handleClick}>
        <div className="file-row-icon">{getFileIcon()}</div>
        <div className="file-row-name">{file.title || file.filename}</div>
        <div className="file-row-date">{formatDate(file.created_at)}</div>
        <div className="file-row-size">{formatSize(file.size)}</div>
        <button className="file-row-action" onClick={(e) => { e.stopPropagation(); onDownload(file); }}>
          <svg viewBox="0 0 24 24" width="18" height="18">
            <path fill="#5f6368" d="M19 9h-4V3H9v6H5l7 7 7-7zM5 18v2h14v-2H5z"/>
          </svg>
        </button>
        <button className="file-row-action" onClick={(e) => { e.stopPropagation(); onDelete(file.id); }}>
          <svg viewBox="0 0 24 24" width="18" height="18">
            <path fill="#5f6368" d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/>
          </svg>
        </button>
      </div>
    );
  }

  return (
    <div className="file-card" onClick={handleClick}>
      <div className="file-thumbnail" style={{ backgroundColor: getThumbnailColor() }}>
        {isVideo && (
          <div className="play-overlay">
            <svg viewBox="0 0 24 24" width="48" height="48">
              <path fill="rgba(0,0,0,0.7)" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 14.5v-9l6 4.5-6 4.5z"/>
            </svg>
          </div>
        )}
        <div className="file-icon-large">
          {isVideo ? (
            <svg viewBox="0 0 24 24" width="48" height="48">
              <path fill="#ea4335" d="M18 4l2 4h-3l-2-4h-2l2 4h-3l-2-4H8l2 4H7L5 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4h-4z"/>
            </svg>
          ) : isImage ? (
            <svg viewBox="0 0 24 24" width="48" height="48">
              <path fill="#4285f4" d="M21 19V5c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zM8.5 13.5l2.5 3.01L14.5 12l4.5 6H5l3.5-4.5z"/>
            </svg>
          ) : isPdf ? (
            <svg viewBox="0 0 24 24" width="48" height="48">
              <path fill="#ea4335" d="M20 2H8c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-8.5 7.5c0 .83-.67 1.5-1.5 1.5H9v2H7.5V7H10c.83 0 1.5.67 1.5 1.5v1zm5 2c0 .83-.67 1.5-1.5 1.5h-2.5V7H15c.83 0 1.5.67 1.5 1.5v3zm4-3H19v1h1.5V11H19v2h-1.5V7h3v1.5zM9 9.5h1v-1H9v1zM4 6H2v14c0 1.1.9 2 2 2h14v-2H4V6zm10 5.5h1v-3h-1v3z"/>
            </svg>
          ) : (
            <svg viewBox="0 0 24 24" width="48" height="48">
              <path fill="#5f6368" d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zm2 16H8v-2h8v2zm0-4H8v-2h8v2zm-3-5V3.5L18.5 9H13z"/>
            </svg>
          )}
        </div>
      </div>
      <div className="file-info">
        <div className="file-name-row">
          <span className="file-icon-small">{getFileIcon()}</span>
          <span className="file-name">{file.title || file.filename}</span>
        </div>
        <div className="file-meta">
          <span>{formatDate(file.created_at)}</span>
        </div>
      </div>
      <div className="file-actions">
        <button className="file-action-btn" title="Download" onClick={(e) => { e.stopPropagation(); onDownload(file); }}>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#5f6368">
            <path d="M19 9h-4V3H9v6H5l7 7 7-7zM5 18v2h14v-2H5z"/>
          </svg>
        </button>
        <button className="file-action-btn" title="Delete" onClick={(e) => { e.stopPropagation(); onDelete(file.id); }}>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#5f6368">
            <path d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/>
          </svg>
        </button>
      </div>
    </div>
  );
}
