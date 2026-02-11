import Head from 'next/head';
import Link from 'next/link';
import Header from '../components/Header';
import { useCart } from '../context/CartContext';
import { FaTrash } from 'react-icons/fa';

export default function CartPage() {
  const { items, removeItem, updateQuantity, totalPrice, totalItems } = useCart();

  if (items.length === 0) {
    return (
      <>
        <Head>
          <title>Корзина – Medical Farm</title>
        </Head>
        <Header />
        <main className="container" style={{ textAlign: 'center', paddingTop: '60px' }}>
          <h1 style={{ fontSize: '2rem', marginBottom: '20px' }}>Корзина пуста</h1>
          <p style={{ marginBottom: '30px' }}>Добавьте товары из каталога</p>
          <Link href="/" className="btn">
            Перейти в каталог
          </Link>
        </main>
      </>
    );
  }

  return (
    <>
      <Head>
        <title>Корзина – Medical Farm</title>
      </Head>
      <Header />
      <main className="container" style={{ paddingTop: '30px', paddingBottom: '60px' }}>
        <h1 style={{ fontSize: '2rem', marginBottom: '30px' }}>Корзина</h1>
        <div style={{ display: 'grid', gridTemplateColumns: '2fr 1fr', gap: '30px' }}>
          <div>
            {items.map((item) => (
              <div
                key={item.inventory_id}
                style={{
                  display: 'flex',
                  justifyContent: 'space-between',
                  alignItems: 'center',
                  padding: '20px',
                  backgroundColor: 'white',
                  borderRadius: '8px',
                  marginBottom: '12px',
                  border: '1px solid #e2e8f0',
                }}
              >
                <div>
                  <h3 style={{ marginBottom: '6px' }}>{item.name}</h3>
                  <p style={{ color: '#475569' }}>{item.price.toFixed(2)} ₽</p>
                </div>
                <div style={{ display: 'flex', alignItems: 'center', gap: '20px' }}>
                  <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                    <button
                      onClick={() => updateQuantity(item.inventory_id, item.quantity - 1)}
                      style={{ padding: '4px 10px', fontSize: '16px' }}
                    >
                      -
                    </button>
                    <span style={{ minWidth: '30px', textAlign: 'center' }}>{item.quantity}</span>
                    <button
                      onClick={() => updateQuantity(item.inventory_id, item.quantity + 1)}
                      style={{ padding: '4px 10px', fontSize: '16px' }}
                    >
                      +
                    </button>
                  </div>
                  <span style={{ fontWeight: 'bold' }}>
                    {(item.price * item.quantity).toFixed(2)} ₽
                  </span>
                  <button
                    onClick={() => removeItem(item.inventory_id)}
                    style={{ background: 'none', border: 'none', color: '#ef4444', cursor: 'pointer' }}
                  >
                    <FaTrash />
                  </button>
                </div>
              </div>
            ))}
          </div>
          <div
            style={{
              backgroundColor: 'white',
              padding: '24px',
              borderRadius: '12px',
              border: '1px solid #e2e8f0',
              height: 'fit-content',
            }}
          >
            <h2 style={{ fontSize: '1.5rem', marginBottom: '20px' }}>Итого</h2>
            <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '12px' }}>
              <span>Товаров:</span>
              <span>{totalItems} шт.</span>
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '24px' }}>
              <span>Сумма:</span>
              <span style={{ fontSize: '1.4rem', fontWeight: 'bold' }}>
                {totalPrice.toFixed(2)} ₽
              </span>
            </div>
            <Link href="/checkout" className="btn" style={{ width: '100%', textAlign: 'center' }}>
              Оформить заказ
            </Link>
          </div>
        </div>
      </main>
    </>
  );
}