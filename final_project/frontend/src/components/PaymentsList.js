import React, { useEffect, useState } from "react";
import { fetchPayments, deletePayment } from "../services/payments";

const PaymentsList = () => {
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    const loadPayments = async () => {
      const data = await fetchPayments();
      setPayments(data);
    };
    loadPayments();
  }, []);

  const handleDelete = async (id) => {
    await deletePayment(id);
    setPayments(payments.filter((payment) => payment.payment_id !== id));
  };

  return (
    <div>
      <h1>Payments</h1>
      <ul>
        {payments.map((payment) => (
          <li key={payment.payment_id}>
            Amount: {payment.amount}, Method: {payment.payment_method}{" "}
            <button onClick={() => handleDelete(payment.payment_id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PaymentsList;
