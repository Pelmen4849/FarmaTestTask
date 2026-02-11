import Link from 'next/link';
import Head from 'next/head';
import Header from '../components/Header';

export default function OrderSuccessPage() {
  return (
    <>
      <Head>
        <title>Заказ оформлен – Medical Farm</title>
      </Head>
      <Header />
      <main className="container" style={{ textAlign: 'center', paddingTop: '60px' }}>
        <h1 style={{ fontSize: '2.5rem', marginBottom: '20px' }}>🎉 Заказ оформлен!</h1>
        <p style={{ fontSize: '1.2rem', marginBottom: '30px' }}>
          Спасибо за покупку. Мы свяжемся с вами для уточнения деталей.
        </p>
        <Link href="/" className="btn">
          Вернуться на главную
        </Link>
      </main>
    </>
  );
}