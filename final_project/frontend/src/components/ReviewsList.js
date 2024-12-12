import React, { useEffect, useState } from "react";
import { fetchReviews, deleteReview } from "../services/reviews";

const ReviewsList = () => {
  const [reviews, setReviews] = useState([]);

  useEffect(() => {
    const loadReviews = async () => {
      const data = await fetchReviews();
      setReviews(data);
    };
    loadReviews();
  }, []);

  const handleDelete = async (id) => {
    await deleteReview(id);
    setReviews(reviews.filter((review) => review.review_id !== id));
  };

  return (
    <div>
      <h1>Reviews</h1>
      <ul>
        {reviews.map((review) => (
          <li key={review.review_id}>
            Product: {review.product_id}, Rating: {review.rating}{" "}
            <button onClick={() => handleDelete(review.review_id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ReviewsList;
