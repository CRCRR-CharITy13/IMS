export interface SignupCode {
    ID: number;
    code: string;
    expired: boolean;
    designatedUsername: string;
    createdByUserID: number;
    createdAt: number;
    expiration: number;
}