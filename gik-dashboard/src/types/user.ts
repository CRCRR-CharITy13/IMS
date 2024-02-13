export interface User {
    ID: number;
    username: string;
    registeredAt: number;
    admin: boolean;
    disabled: boolean;
}

export interface SimpleUser {
    ID: string;
    username: string;
}