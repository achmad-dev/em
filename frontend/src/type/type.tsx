/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

interface Event {
    id: string;
    user_id: string;
    vendor_name: string;
    event_name: string;
    proposed_dates: Date[];
    rejected_remarks?: string;
    status: string;
    confirmed_date?: Date;
    postal_code: string;
    location: string;
    created_at: Date;
    updated_at: Date;
};

interface CompanyRes {
    company_name: string;
}

interface RejectRequest {
    event_id: string;
    rejected_remarks: string;
    status: string;
}

interface ConfirmRequest {
    event_id: string;
    confirmed_date: Date;
    status: string;
}

type AuthResponse = {
    token: string;
    role: string;
}

interface CreateEventRequest {
    name: string;
    vendor_name: string;
    postal_code: string;
    location: string;
    proposed_dates: Date[];
}

interface EventVendorOption {
    event_id: string;
    remarks?: string;
    confirmed_date?: Date;
    status: string;
}

export type { Event, CompanyRes, RejectRequest, ConfirmRequest, AuthResponse, CreateEventRequest, EventVendorOption };