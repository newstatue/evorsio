import {
  File,
  FileArchive,
  FileAudio,
  FileCode,
  FileImage,
  FileJson,
  FileSpreadsheet,
  FileText,
  FileVideo,
  Folder,
  type LucideIcon,
} from "lucide-react"
import { EntryType } from "@bindings/github.com/newstatue/evorsio/internal/drive"

export function getDriveIcon(mime?: string, type?: EntryType): LucideIcon {
  if (type === "folder") {
    return Folder
  }

  if (!mime) {
    return File
  }

  if (mime.startsWith("image/")) {
    return FileImage
  }

  if (mime.startsWith("video/")) {
    return FileVideo
  }

  if (mime.startsWith("audio/")) {
    return FileAudio
  }

  switch (mime) {
    case "application/pdf":
    case "text/plain":
    case "text/markdown":
    case "application/msword":
    case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
      return FileText

    case "application/json":
      return FileJson

    case "text/javascript":
    case "application/javascript":
    case "text/typescript":
    case "application/typescript":
    case "text/html":
    case "text/css":
      return FileCode

    case "text/csv":
    case "application/vnd.ms-excel":
    case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
      return FileSpreadsheet

    case "application/zip":
    case "application/x-rar-compressed":
    case "application/x-7z-compressed":
    case "application/gzip":
    case "application/x-tar":
      return FileArchive

    default:
      return File
  }
}
