/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import type { Event } from "@/type/type";

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
import { Label } from "@/components/ui/label";
import { PopUpRejectEvent } from "./popUpRejectEvent";
import { PopUpAcceptEvent } from "./popUpAcceptEvent";

const PopUpViewEvent = (prop: { event: Event, role: string }) => {
  const { event } = prop;
  console.log(prop.role);

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>View Event</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Event Details</DialogTitle>
          <DialogDescription>
            Here are the details for the event.
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">Name</Label>
            <p className="col-span-3">{event.event_name}</p>
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">Vendor Name</Label>
            <p className="col-span-3">{event.vendor_name}</p>
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">Postal Code</Label>
            <p className="col-span-3">{event.postal_code}</p>
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">Location</Label>
            <p className="col-span-3">{event.location}</p>
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">Proposed Dates</Label>
            <div className="col-span-3 grid gap-2">
              {event.proposed_dates.map((date, index) => (
                <p key={index}>{new Date(date).toLocaleDateString()}</p>
              ))}
            </div>
          </div>
          {event.confirmed_date && (
            <div className="grid grid-cols-4 items-center gap-4">
              <Label className="text-right">Confirmed Date</Label>
              <p className="col-span-3">
                {new Date(event.confirmed_date).toLocaleDateString()}
              </p>
            </div>
          )}
          {event.rejected_remarks && (
            <div className="grid grid-cols-4 items-center gap-4">
              <Label className="text-right">Rejected Remarks</Label>
              <p className="col-span-3">{event.rejected_remarks}</p>
            </div>
          )}
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">Status</Label>
            <p className="col-span-3">{event.status}</p>
          </div>
        </div>
        <DialogFooter>
            {prop.role === "vendor" && (
                <>
                    <PopUpRejectEvent event={event} />
                    <PopUpAcceptEvent event={event} />
                </>
            )}
          <DialogTrigger asChild>
            <Button type="button">Close</Button>
          </DialogTrigger>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default PopUpViewEvent;
