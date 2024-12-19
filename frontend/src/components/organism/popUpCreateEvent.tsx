/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { CreateEventRequest } from "@/type/type";
import { Label } from "@/components/ui/label";
import type { CompanyRes } from "@/type/type";
import { useEvent } from "@/hooks/useEvent";

export function DialogCreateEvent() {
  const { getVendors, createEvent } = useEvent();

  const [companies, setCompanies] = useState<CompanyRes[]>([]);
  const [formData, setFormData] = useState({
    name: "",
    vendorName: "",
    postalCode: "",
    location: "",
    proposedDates: ["", "", ""],
  });

  useEffect(() => {
    async function fetchVendors() {
      const vendors = await getVendors();
      setFormData((prev) => ({
        ...prev,
        vendorName: vendors[0].company_name,
      }));
      setCompanies(vendors);
    }
    if (companies.length === 0) {
      fetchVendors();
    }
  }, [companies.length, getVendors]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { id, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [id]: value,
    }));
  };

  const handleDateChange = (index: number, value: string) => {
    const newDates = [...formData.proposedDates];
    newDates[index] = value;
    setFormData((prev) => ({
      ...prev,
      proposedDates: newDates,
    }));
  };

  const handleSubmit = async () => {
    const formattedDates = formData.proposedDates.map((date) => new Date(date));

    const createEventRequest: CreateEventRequest = {
      name: formData.name,
      vendor_name: formData.vendorName,
      postal_code: formData.postalCode,
      location: formData.location,
      proposed_dates: formattedDates,
    };

    await createEvent(createEventRequest);
    setFormData({
      name: "",
      vendorName: "",
      postalCode: "",
      location: "",
      proposedDates: ["", "", ""],
    });
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Create Event</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create Event</DialogTitle>
          <DialogDescription>
            Fill in the details for the event. Click save when you're done.
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input
              id="name"
              placeholder="Event Name"
              className="col-span-3"
              value={formData.name}
              onChange={handleChange}
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="vendorName" className="text-right">
              Vendor Name
            </Label>
            <select
              id="vendorName"
              className="col-span-3 border rounded p-2"
              value={formData.vendorName}
              onChange={handleChange}
            >
              {companies.map((company, index) => (
                <option key={index} value={company.company_name}>
                  {company.company_name}
                </option>
              ))}
            </select>
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="postalCode" className="text-right">
              Postal Code
            </Label>
            <Input
              id="postalCode"
              placeholder="Postal Code"
              className="col-span-3"
              value={formData.postalCode}
              onChange={handleChange}
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="location" className="text-right">
              Location
            </Label>
            <Input
              id="location"
              placeholder="Location"
              className="col-span-3"
              value={formData.location}
              onChange={handleChange}
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="proposedDates" className="text-right">
              Proposed Dates
            </Label>
            <div className="col-span-3 grid gap-2">
              <Input
                id="proposedDate1"
                type="date"
                value={formData.proposedDates[0]}
                onChange={(e) => handleDateChange(0, e.target.value)}
              />
              <Input
                id="proposedDate2"
                type="date"
                value={formData.proposedDates[1]}
                onChange={(e) => handleDateChange(1, e.target.value)}
              />
              <Input
                id="proposedDate3"
                type="date"
                value={formData.proposedDates[2]}
                onChange={(e) => handleDateChange(2, e.target.value)}
              />
            </div>
          </div>
        </div>
        <DialogFooter>
          <DialogTrigger asChild>
            <Button type="button" onClick={handleSubmit}>
              Save changes
            </Button>
          </DialogTrigger>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
