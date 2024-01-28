import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Home = () => {
  const [amount, setAmount] = useState(0);
  const navigate = useNavigate(); // Hook to navigate to different routes

  const handleInputChange = (e) => {
    setAmount(e.target.value);
  };

  const handlePay = () => {
    // Add your payment logic here
    console.log(`Payment of $${amount} processed.`);
    navigate('/card'); // Navigate to the CardDetailsPage upon payment
  };

  return (
    <div className="max-w-xs mx-auto bg-white rounded p-4 shadow-md">
      <h2 className="text-lg font-semibold mb-4">Payment Box</h2>
      <input
        type="number"
        placeholder="Enter amount in dollars"
        value={amount}
        onChange={handleInputChange}
        className="w-full px-3 py-2 border rounded mb-4"
      />
      <button
        onClick={handlePay}
        className="w-full bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600"
      >
        Pay
      </button>
    </div>
  );
};

export default Home;
