export interface Image {
    base64: string;
    extension: string;
    fileName: string;
}

export interface ImageInfo {
    ID: string
    Name: string
    Size: number
    Dimensions: string
    FileTypes: FileType[]
    Extension: string
    Alt: string
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: any
}

export interface FileType {
    ID: string
    Name: string
    UpdatedAt: string
    CreatedAt: string
    DeletedAt: any
}
