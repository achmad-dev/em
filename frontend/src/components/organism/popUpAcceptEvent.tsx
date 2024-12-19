/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import type { Event } from "@/type/type";
import { useState } from "react";
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
import type { ConfirmRequest } from "@/type/type";
import { useEvent } from "@/hooks/useEvent";


export const PopUpAcceptEvent = (prop: { event: Event }) => {
    const [confirmedDate, setConfirmedDate] = useState<Date | null>(null);
    const { updateEvent } = useEvent();

    const handleAccept = () => {
        if (confirmedDate) {
            const confirmRequest: ConfirmRequest = {
                event_id: prop.event.id,
                confirmed_date: confirmedDate,
                status: "confirmed",
            };
            // Handle the confirm request here
            updateEvent(confirmRequest);
        }
    };

    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button className="bg-green-500 text-white">Accept Event</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Accept Event</DialogTitle>
                    <DialogDescription>
                        Select a proposed date to confirm the event. Click save when you're done.
                    </DialogDescription>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="confirmedDate" className="text-right">
                            Confirm Date
                        </Label>
                        <select id="vendorName" onChange={(e) => {
                            setConfirmedDate(new Date(e.target.value));
                        }} className="col-span-3 border rounded p-2">
                            {prop.event.proposed_dates.map((date, index) => (
                                <option key={index} value={date.toString()}>
                                    {date.toString()}
                                </option>
                            ))}
                        </select>
                    </div>
                </div>
                <DialogFooter>
                    <DialogTrigger asChild>
                        <Button type="button" onClick={handleAccept}>
                            Save changes
                        </Button>
                    </DialogTrigger>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
};