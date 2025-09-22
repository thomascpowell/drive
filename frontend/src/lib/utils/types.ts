export interface Res {
  message?: string;
  error?: string;
  token?: string
}

export interface FileRec {
  ID: number
  Filename: string
  Path: string
  Size: number
  UploadedAt: string
  UploadedBy: number
}

export interface Share {
  FileID: number
  TTL: number
}

