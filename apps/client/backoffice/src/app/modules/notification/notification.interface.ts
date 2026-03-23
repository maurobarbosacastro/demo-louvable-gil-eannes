
export interface NotificationMessage {
    id: string;
    title: string;
    body: string;
    duration?: number;
    dismissible?: boolean;
}
