export interface Option {
    label: string[];
    value: string;
    icons?: string[]
    image?: Image;
}

export interface Image {
    base64: string;
    extension: string;
    fileName: string;
}
