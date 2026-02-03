import { useRef, useEffect } from 'react';

interface VideoPlayerProps {
  url: string;
  title?: string;
  onClose: () => void;
}

export function VideoPlayer({ url, title, onClose }: VideoPlayerProps) {
  const videoRef = useRef<HTMLVideoElement>(null);

  useEffect(() => {
    const handleEscape = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose();
    };
    window.addEventListener('keydown', handleEscape);
    return () => window.removeEventListener('keydown', handleEscape);
  }, [onClose]);

  return (
    <div className="video-overlay" onClick={onClose}>
      <div className="video-container" onClick={(e) => e.stopPropagation()}>
        <button className="close-btn" onClick={onClose}>×</button>
        {title && <h3 className="video-title">{title}</h3>}
        <video
          ref={videoRef}
          src={url}
          controls
          autoPlay
          className="video-player"
        />
      </div>
    </div>
  );
}
