import { useState } from 'react';
import { useRouter } from 'next/router';
import Head from 'next/head';
import Header from '../components/Header';
import { useCart } from '../context/CartContext';
import { createOrder } from '../utils/api';

export default function CheckoutPage() {
  const router = useRouter();
  const { items, totalPrice, clearCart } = useCart();
  const [customer, setCustomer] = useState({
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
  });
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    // В реальном приложении здесь нужно сначала создать/получить customer_id
    // Пока используем заглушку customer_id = 1
    const orderData = {
      customer_id: 1,
      shop_id: 1,
      items: items.map((item) => ({
        inventory_id: item.inventory_id,
        quantity: item.quantity,
      })),
    };

    try {
      await createOrder(orderData);
      clearCart();
      router.push('/order-success');
    } catch (error) {
      console.error('Order failed:', error);
      alert('Ошибка при оформлении заказа. Попробуйте позже.');
    } finally {
      setIsSubmitting(false);
    }
  };

  if (items.length === 0) {
    return (
      <>
        <Head>
          <title>Оформление заказа – Medical Farm</title>
        </Head>
        <Header />
        <main className="container" style={{ textAlign: 'center', paddingTop: '60px' }}>
          <h1>Корзина пуста</h1>
          <p style={{ marginBottom: '20px' }}>Добавьте товары для оформления заказа</p>
          <button onClick={() => router.push('/')} className="btn">
            Перейти в каталог
          </button>
        </main>
      </>
    );
  }

  return (
    <>
      <Head>
        <title>Оформление заказа – Medical Farm</title>
      </Head>
      <Header />
      <main className="container" style={{ paddingTop: '30px', paddingBottom: '60px' }}>
        <h1 style={{ fontSize: '2rem', marginBottom: '30px' }}>Оформление заказа</h1>
        <div style={{ display: 'grid', gridTemplateColumns: '2fr 1fr', gap: '30px' }}>
          <div
            style={{
              backgroundColor: 'white',
              padding: '30px',
              borderRadius: '12px',
              border: '1px solid #e2e8f0',
            }}
          >
            <form onSubmit={handleSubmit}>
              <h2 style={{ marginBottom: '20px' }}>Контактные данные</h2>
              <div style={{ marginBottom: '16px' }}>
                <label htmlFor="firstName">Имя</label>
                <input
                  id="firstName"
                  type="text"
                  required
                  value={customer.firstName}
                  onChange={(e) => setCustomer({ ...customer, firstName: e.target.value })}
                />
              </div>
              <div style={{ marginBottom: '16px' }}>
                <label htmlFor="lastName">Фамилия</label>
                <input
                  id="lastName"
                  type="text"
                  required
                  value={customer.lastName}
                  onChange={(e) => setCustomer({ ...customer, lastName: e.target.value })}
                />
              </div>
              <div style={{ marginBottom: '16px' }}>
                <label htmlFor="email">Email</label>
                <input
                  id="email"
                  type="email"
                  required
                  value={customer.email}
                  onChange={(e) => setCustomer({ ...customer, email: e.target.value })}
                />
              </div>
              <div style={{ marginBottom: '24px' }}>
                <label htmlFor="phone">Телефон</label>
                <input
                  id="phone"
                  type="tel"
                  required
                  value={customer.phone}
                  onChange={(e) => setCustomer({ ...customer, phone: e.target.value })}
                />
              </div>
              <button
                type="submit"
                className="btn"
                disabled={isSubmitting}
                style={{ width: '100%' }}
              >
                {isSubmitting ? 'Оформляем...' : 'Подтвердить заказ'}
              </button>
            </form>
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
            <h2 style={{ fontSize: '1.5rem', marginBottom: '20px' }}>Ваш заказ</h2>
            {items.map((item) => (
              <div
                key={item.inventory_id}
                style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '12px' }}
              >
                <span>
                  {item.name} x{item.quantity}
                </span>
                <span>{(item.price * item.quantity).toFixed(2)} ₽</span>
              </div>
            ))}
            <div
              style={{
                borderTop: '1px solid #e2e8f0',
                marginTop: '16px',
                paddingTop: '16px',
              }}
            >
              <div style={{ display: 'flex', justifyContent: 'space-between', fontWeight: 'bold' }}>
                <span>Итого:</span>
                <span>{totalPrice.toFixed(2)} ₽</span>
              </div>
            </div>
          </div>
        </div>
      </main>
    </>
  );
}