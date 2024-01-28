import React, { useState } from 'react';
import { Box, Button, HStack, Text, useToast } from '@chakra-ui/react';
import { useStripe, useElements, PaymentElement } from '@stripe/react-stripe-js';

import { useRouter } from 'next/router';
import { useAuthContext } from '../../../context/AuthContext';
import { usePaymentContext } from '../../../context/PaymentContext';
import useRoleName from '../../../hooks/useRoleName';

const CardDetailsPage = () => {
  const { paseto } = useAuthContext();
  const { stripePayload } = usePaymentContext();
  const roleName = useRoleName();

  const stripe = useStripe();
  const elements = useElements();
  const toast = useToast();
  const router = useRouter();

  const [transactionError, setTransactionError] = useState('');
  const [transactionStatus, setTransactionStatus] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (!stripe || !elements) {
      return;
    }

    const result = await stripe.confirmPayment({
      elements,
      confirmParams: {
        return_url: `${process.env.NEXT_PUBLIC_ORIGIN}/dats`,
      },
    });

    if (result.error) {
      setTransactionError(result.error.message);
      setTransactionStatus('rejected');
      toast({
        title: result.error.message,
        status: 'error',
        isClosable: true,
      });
    } else {
      setTransactionStatus('approved');
      toast({
        title: 'Transaction has been approved',
        status: 'success',
        isClosable: true,
      });
      // Redirect or perform additional logic upon successful payment
    }
  };

  return (
    <Box w="100%" mt="9" maxW="500px">
      {!elements ? <Text>Loading form...</Text> : null}
      <form id="payment-form" onSubmit={handleSubmit}>
        <PaymentElement id="payment-element" />
        <HStack mt="5" spacing={3}>
          <Button colorScheme="brand" type="submit">
            Pay
          </Button>
        </HStack>
      </form>
    </Box>
  );
};

export default CardDetailsPage;
