export interface Donation {
    ID: number;
    createdTime: string;
    donorId: number;
    donorBy: string;
    signerId : number;
    signedBy: string;
    totalValue: number;
}

export interface DonationItem {
    ID: number;
    name: string;
    sku: string;
    size: string;
    price: number;
    quantity: number;
    totalValue: number;
}

export interface AddDonationResponse {
    ID: number;
    itemSKUName : string;
    locationName : string;
    quantity : number;
}