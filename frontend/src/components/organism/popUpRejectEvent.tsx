/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import type { Event, RejectRequest } from "@/type/type";
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
import { Textarea } from "@/components/ui/textarea";
import { useState } from "react";
import { Label } from "../ui/label";
import { useEvent } from "@/hooks/useEvent";

export const PopUpRejectEvent = (prop: { event: Event }) => {
  const [rejectedRemarks, setRejectedRemarks] = useState("");
  const { updateEvent } = useEvent();

  const handleReject = () => {
    const rejectRequest: RejectRequest = {
      event_id: prop.event.id,
      rejected_remarks: rejectedRemarks,
      status: "rejected",
    };
    // Handle the reject request here
    updateEvent(rejectRequest);
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="bg-red-500 text-white">Reject Event</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Reject Event</DialogTitle>
          <DialogDescription>
            Provide remarks for rejecting the event. Click save when you're
            done.
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="rejectedRemarks" className="text-right">
              Remarks
            </Label>
            <Textarea
              id="rejectedRemarks"
              placeholder="Enter rejection remarks"
              className="col-span-3"
              value={rejectedRemarks}
              onChange={(e) => setRejectedRemarks(e.target.value)}
            />
          </div>
        </div>
        <DialogFooter>
          <DialogTrigger asChild>
            <Button type="button" onClick={handleReject}>
              Save changes
            </Button>
          </DialogTrigger>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
