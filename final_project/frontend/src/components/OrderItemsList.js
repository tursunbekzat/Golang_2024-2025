import React, { useEffect, useState } from "react";
import { fetchOrderItems, deleteOrderItem } from "../services/orderItems";

const OrderItemsList = () => {
  const [orderItems, setOrderItems] = useState([]);

  useEffect(() => {
    const loadOrderItems = async () => {
      const data = await fetchOrderItems();
      setOrderItems(data);
    };
    loadOrderItems();
  }, []);

  const handleDelete = async (id) => {
    await deleteOrderItem(id);
    setOrderItems(orderItems.filter((item) => item.order_item_id !== id));
  };

  return (
    <div>
      <h1>Order Items</h1>
      <ul>
        {orderItems.map((item) => (
          <li key={item.order_item_id}>
            Product: {item.product_id}, Quantity: {item.quantity}{" "}
            <button onClick={() => handleDelete(item.order_item_id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default OrderItemsList;
