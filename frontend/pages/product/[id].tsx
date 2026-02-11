import { GetServerSideProps } from 'next';
import { useRouter } from 'next/router';
import { useState } from 'react';
import Head from 'next/head';
import Image from 'next/image';
import Header from '../../components/Header';
import { useCart } from '../../context/CartContext';
import { getDrugById, getDrugs, Drug, InventoryItem } from '../../utils/api';

interface ProductPageProps {
  drug: Drug;
  inventory: InventoryItem[];
}

export default function ProductPage({ drug, inventory }: ProductPageProps) {
  const router = useRouter();
  const { addItem } = useCart();
  const [selectedInventory, setSelectedInventory] = useState<InventoryItem | null>(
    inventory[0] || null
  );
  const [quantity, setQuantity] = useState(1);

  const availableStock = selectedInventory?.quantity || 0;
  const price = selectedInventory?.selling_price || 0;

  const handleAddToCart = () => {
    if (!selectedInventory) return;
    addItem({
      inventory_id: selectedInventory.id,
      drug_id: drug.id,
      name: drug.name,
      price: selectedInventory.selling_price,
      quantity,
    });
    alert('Товар добавлен в корзину');
  };

  if (!drug) return <p>Товар не найден</p>;

  return (
    <>
      <Head>
        <title>{drug.name} – Medical Farm</title>
      </Head>
      <Header />
      <main className="container" style={{ paddingTop: '30px', paddingBottom: '60px' }}>
        <button
          onClick={() => router.back()}
          style={{
            background: 'none',
            border: 'none',
            color: '#2563eb',
            cursor: 'pointer',
            marginBottom: '20px',
          }}
        >
          ← Назад
        </button>
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '40px' }}>
          {/* Изображение товара */}
          <div style={{ position: 'relative', width: '100%', height: '400px' }}>
            <Image
              src={`/images/drugs/${drug.id}.jpg`}
              alt={drug.name}
              fill
              style={{ objectFit: 'contain' }}
              sizes="(max-width: 768px) 100vw, 600px"
              onError={(e) => {
                (e.target as HTMLImageElement).src = '/images/placeholder.png';
              }}
            />
          </div>
          {/* Информация о товаре */}
          <div>
            <h1 style={{ fontSize: '2.2rem', marginBottom: '16px' }}>{drug.name}</h1>
            <div style={{ marginBottom: '20px' }}>
              <span
                style={{
                  backgroundColor: drug.requires_prescription ? '#fee2e2' : '#dcfce7',
                  color: drug.requires_prescription ? '#b91c1c' : '#166534',
                  padding: '4px 12px',
                  borderRadius: '20px',
                  fontSize: '0.9rem',
                }}
              >
                {drug.requires_prescription ? 'Требуется рецепт' : 'Без рецепта'}
              </span>
            </div>
            {drug.dosage_form && (
              <p><strong>Форма выпуска:</strong> {drug.dosage_form}</p>
            )}
            {drug.dosage && (
              <p><strong>Дозировка:</strong> {drug.dosage}</p>
            )}
            {drug.description && (
              <p style={{ marginTop: '20px', lineHeight: '1.6' }}>{drug.description}</p>
            )}
          </div>
        </div>

        {/* Блок покупки */}
        <div style={{ marginTop: '40px', backgroundColor: 'white', padding: '30px', borderRadius: '12px', border: '1px solid #e2e8f0' }}>
          {inventory.length === 0 ? (
            <p>Нет в наличии</p>
          ) : (
            <>
              {inventory.length > 1 && (
                <div style={{ marginBottom: '20px' }}>
                  <label>Выберите предложение</label>
                  <select
                    value={selectedInventory?.id}
                    onChange={(e) => {
                      const inv = inventory.find((i) => i.id === Number(e.target.value));
                      setSelectedInventory(inv || null);
                    }}
                  >
                    {inventory.map((inv) => (
                      <option key={inv.id} value={inv.id}>
                        {inv.selling_price} ₽ / {inv.quantity} шт.
                      </option>
                    ))}
                  </select>
                </div>
              )}
              <div style={{ marginBottom: '20px' }}>
                <label>Цена</label>
                <p style={{ fontSize: '2rem', fontWeight: 'bold', color: '#0f172a' }}>
                  {price.toFixed(2)} ₽
                </p>
              </div>
              <div style={{ marginBottom: '30px' }}>
                <label>Количество</label>
                <div style={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
                  <input
                    type="number"
                    min="1"
                    max={availableStock}
                    value={quantity}
                    onChange={(e) => setQuantity(parseInt(e.target.value) || 1)}
                    style={{ width: '80px' }}
                  />
                  <span>доступно: {availableStock}</span>
                </div>
              </div>
              <button
                onClick={handleAddToCart}
                disabled={availableStock === 0}
                className="btn"
                style={{ width: '100%' }}
              >
                Добавить в корзину
              </button>
            </>
          )}
        </div>
      </main>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ params }) => {
  const id = Number(params?.id);
  if (isNaN(id)) return { notFound: true };

  try {
    const drugRes = await getDrugById(id);
    const drug = drugRes.data;
    const inventoryRes = await getDrugs(1);
    const inventory = inventoryRes.data.filter(
      (item) => item.drug_id === id && item.quantity > 0
    );
    return { props: { drug, inventory } };
  } catch (error) {
    console.error(error);
    return { notFound: true };
  }
};