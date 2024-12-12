import React, { useEffect, useState } from "react";
import { fetchShoppingCarts, deleteShoppingCart } from "../services/shoppingCart";

const ShoppingCartsList = () => {
  const [carts, setCarts] = useState([]);

  useEffect(() => {
    const loadCarts = async () => {
      const data = await fetchShoppingCarts();
      setCarts(data);
    };
    loadCarts();
  }, []);

  const handleDelete = async (id) => {
    await deleteShoppingCart(id);
    setCarts(carts.filter((cart) => cart.cart_id !== id));
  };

  return (
    <div>
      <h1>Shopping Carts</h1>
      <ul>
        {carts.map((cart) => (
          <li key={cart.cart_id}>
            User: {cart.user_id}{" "}
            <button onClick={() => handleDelete(cart.cart_id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ShoppingCartsList;
