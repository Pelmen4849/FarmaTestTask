import Image from 'next/image';
import Link from 'next/link';
import { useCart } from '../context/CartContext';

interface ProductCardProps {
  id: number;          
  inventoryId: number; 
  name: string;
  price: number;
  dosageForm?: string | null;
  dosage?: string | null;
  stock: number;
  imageUrl?: string;   
}

export default function ProductCard({
  id,
  inventoryId,
  name,
  price,
  dosageForm,
  dosage,
  stock,
  imageUrl = `/images/drugs/${id}.jpg`, // значение по умолчанию
}: ProductCardProps) {
  const { addItem } = useCart();

  const handleAddToCart = () => {
    addItem({
      inventory_id: inventoryId,
      drug_id: id,
      name,
      price,
      quantity: 1,
    });
  };

  return (
    <div
      style={{
        border: '1px solid #e2e8f0',
        borderRadius: '12px',
        padding: '20px',
        backgroundColor: 'white',
        boxShadow: '0 1px 3px rgba(0,0,0,0.05)',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      {/* Блок изображения */}
      <div style={{ position: 'relative', width: '100%', height: '180px', marginBottom: '16px' }}>
        <Image
          src={imageUrl}
          alt={name}
          fill
          style={{ objectFit: 'contain' }}
          sizes="(max-width: 768px) 100vw, 280px"
          onError={(e) => {
            // Если изображение не загрузилось – показываем заглушку
            (e.target as HTMLImageElement).src = '/images/placeholder.png';
          }}
        />
      </div>

      <h3 style={{ marginBottom: '10px', fontSize: '1.2rem' }}>{name}</h3>
      {(dosageForm || dosage) && (
        <p style={{ color: '#64748b', marginBottom: '8px' }}>
          {dosageForm} {dosage}
        </p>
      )}
      <div style={{ marginTop: 'auto' }}>
        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            marginBottom: '16px',
          }}
        >
          <span style={{ fontWeight: 'bold', fontSize: '1.4rem', color: '#0f172a' }}>
            {price.toFixed(2)} ₽
          </span>
          <span
            style={{
              color: stock > 5 ? '#10b981' : '#f59e0b',
              fontSize: '0.9rem',
            }}
          >
            {stock > 0 ? `В наличии: ${stock}` : 'Нет в наличии'}
          </span>
        </div>
        <div style={{ display: 'flex', gap: '10px' }}>
          <Link
            href={`/product/${id}`}
            className="btn btn-small"
            style={{ flex: 1, textAlign: 'center' }}
          >
            Подробнее
          </Link>
          <button
            onClick={handleAddToCart}
            disabled={stock <= 0}
            className="btn btn-small"
            style={{
              flex: 1,
              backgroundColor: stock > 0 ? '#2563eb' : '#94a3b8',
              cursor: stock > 0 ? 'pointer' : 'not-allowed',
            }}
          >
            В корзину
          </button>
        </div>
      </div>
    </div>
  );
}