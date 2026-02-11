import { GetServerSideProps } from 'next';
import Head from 'next/head';
import Header from '../components/Header';
import ProductCard from '../components/ProductCard';
import { getDrugs, InventoryItem } from '../utils/api';
import styles from '../styles/Home.module.css';

interface HomeProps {
  products: InventoryItem[];
}

export default function Home({ products }: HomeProps) {
  return (
    <>
      <Head>
        <title>Medical Farm – аптека</title>
        <meta name="description" content="Лекарства с доставкой" />
      </Head>
      <Header />
      <main className="container" style={{ paddingTop: '40px', paddingBottom: '60px' }}>
        <h1 style={{ fontSize: '2rem', marginBottom: '8px' }}>Каталог лекарств</h1>
        <p style={{ color: '#475569', marginBottom: '30px' }}>
          Только сертифицированные препараты
        </p>
        {products.length === 0 ? (
          <p>Нет доступных товаров</p>
        ) : (
          <div className={styles.grid}>
            {products.map((item) => (
              <ProductCard
                key={item.id}
                id={item.drug_id}
                inventoryId={item.id}
                name={item.drug_name || 'Лекарство'}
                price={item.selling_price}
                dosageForm={item.dosage_form}
                dosage={item.dosage}
                stock={item.quantity}
                imageUrl={`/images/drugs/${item.drug_id}.jpg`} 
              />
            ))}
          </div>
        )}
      </main>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async () => {
  try {
    const res = await getDrugs(2);
    const products = res.data;
    return { props: { products } };
  } catch (error) {
    console.error('Failed to fetch products:', error);
    return { props: { products: [] } };
  }
};