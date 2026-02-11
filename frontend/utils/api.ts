import axios from 'axios';

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export interface Drug {
  id: number;
  name: string;
  international_name?: string | null;
  manufacturer_id: number;
  category_id?: number | null;
  description?: string | null;
  dosage_form?: string | null;
  dosage?: string | null;
  requires_prescription: boolean;
  storage_conditions?: string | null;
  expiry_days?: number | null;
  created_at: string;
  updated_at: string;
}

export interface InventoryItem {
  id: number;
  shop_id: number;
  drug_id: number;
  batch_number?: string | null;
  quantity: number;
  purchase_price: number;
  selling_price: number;
  manufacturing_date?: string | null;
  expiry_date: string;
  supplier_id?: number | null;
  received_at: string;
  last_updated: string;
  drug_name?: string;
  dosage_form?: string;
  dosage?: string;
}

export interface OrderRequest {
  customer_id: number;
  shop_id: number;
  items: {
    inventory_id: number;
    quantity: number;
  }[];
}

export interface OrderResponse {
  id: number;
  order_number: string;
  final_amount: number;
  status: string;
}

export const getDrugs = (shopId = 2) =>
  api.get<InventoryItem[]>('/drugs', { params: { shop_id: shopId } });

export const getDrugById = (id: number) => api.get<Drug>(`/drugs/${id}`);

export const createOrder = (order: OrderRequest) =>
  api.post<OrderResponse>('/orders', order);