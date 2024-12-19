/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import useStore from "@/store/useStore";
import instance from "./base";
import type {
  Event,
  CreateEventRequest,
  EventVendorOption,
  CompanyRes,
} from "@/type/type";
import { useQuery, useMutation, useQueryClient } from "react-query";
import { v4 as uuidv4 } from "uuid";
import GenerateSignature from "@/utils/signature";

export const useEvent = () => {
  const queryClient = useQueryClient();
  const { token } = useStore();

  const fetchEvents = async () => {
    const timestamp = new Date().toISOString();
    const key = import.meta.env.VITE_SECRET_KEY_HMAC || "key";
    const signature = GenerateSignature(key, timestamp);

    const response = await instance.get("/events", {
      headers: {
        Authorization: `Bearer ${token}`,
        "X-RequestId": uuidv4(),
        "X-Signature": signature,
        "X-TimeStamp": timestamp,
      },
    });
    const data = response.data.data as Event[];
    return data;
  };

  const createEvent = async (event: CreateEventRequest) => {
    const timestamp = new Date().toISOString();
    const key = import.meta.env.VITE_SECRET_KEY_HMAC || "key";
    const signature = GenerateSignature(key, timestamp);

    return instance.post("/events/insert", event, {
      headers: {
        Authorization: `Bearer ${token}`,
        "X-RequestId": uuidv4(),
        "X-Signature": signature,
        "X-TimeStamp": timestamp,
      },
    });
  };

  const updateEvent = async (event: EventVendorOption) => {
    const timestamp = new Date().toISOString();
    const key = import.meta.env.VITE_SECRET_KEY_HMAC || "key";
    const signature = GenerateSignature(key, timestamp);

    return instance.put(`/events/update?eventId=${event.event_id}`, event, {
      headers: {
        Authorization: `Bearer ${token}`,
        "X-RequestId": uuidv4(),
        "X-Signature": signature,
        "X-TimeStamp": timestamp,
      },
    });
  };

  const getVendors = async () => {
    const timestamp = new Date().toISOString();
    const key = import.meta.env.VITE_SECRET_KEY_HMAC || "key";
    const signature = GenerateSignature(key, timestamp);

    const response = await instance.get("/events/companies", {
      headers: {
        Authorization: `Bearer ${token}`,
        "X-RequestId": uuidv4(),
        "X-Signature": signature,
        "X-TimeStamp": timestamp,
      },
    });
    const data = response.data.data as CompanyRes[];
    return data;
  };

  const {
    data: events,
    isLoading: loading,
    error,
  } = useQuery("events", fetchEvents);

  const createEventMutation = useMutation(createEvent, {
    onSuccess: () => {
      queryClient.invalidateQueries("events");
    },
    onError: () => {
      console.error("Failed to create event");
    },
  });

  const updateEventMutation = useMutation(updateEvent, {
    onSuccess: () => {
      queryClient.invalidateQueries("events");
    },
    onError: () => {
      console.error("Failed to update event");
    },
  });

  return {
    events,
    loading,
    error,
    queryClient,
    getVendors: getVendors,
    createEvent: createEventMutation.mutate,
    updateEvent: updateEventMutation.mutate,
  };
};
