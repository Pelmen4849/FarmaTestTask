import Link from 'next/link';
import { useCart } from '../context/CartContext';
import { FaShoppingCart } from 'react-icons/fa';

export default function Header() {
  const { totalItems } = useCart();

  return (
    <header style={{ backgroundColor: '#0f172a', padding: '1rem 0' }}>
      <div
        className="container"
        style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}
      >
        <Link
          href="/"
          style={{
            color: 'white',
            fontSize: '1.5rem',
            fontWeight: 'bold',
            textDecoration: 'none',
          }}
        >
          ⚕️ Medical Farm
        </Link>
        <Link
          href="/cart"
          style={{
            color: 'white',
            display: 'flex',
            alignItems: 'center',
            gap: '6px',
            textDecoration: 'none',
          }}
        >
          <FaShoppingCart size={20} />
          Корзина
          {totalItems > 0 && (
            <span
              style={{
                backgroundColor: '#ef4444',
                borderRadius: '50%',
                padding: '2px 8px',
                fontSize: '12px',
                fontWeight: 'bold',
              }}
            >
              {totalItems}
            </span>
          )}
        </Link>
      </div>
    </header>
  );
}