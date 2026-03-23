export interface PaginationInterface<T> {
    limit: number;
    page: number;
    sort: string;
    totalPages: number;
    totalRows: number;
    data: T[];
}
