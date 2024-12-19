/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import type { Event } from "@/type/type";
import { cn } from "@/lib/utils";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import PopUpViewEvent from "./popUpViewEvent";

export function EventCard(props: { event: Event; role: string }) {
  const { event } = props;

  return (
    <Card className={cn("w-auto")}>
      <CardHeader>
        <CardTitle>{event.event_name}</CardTitle>
        <CardDescription>{event.vendor_name}</CardDescription>
      </CardHeader>
      {event.confirmed_date ? (
        <CardContent className="flex-1 h-20 max-h-20">
          <CardDescription>
            <p className="text-sm font-medium leading-none">Confirmed Date:</p>
            <p className="text-sm text-muted-foreground">
              {new Date(event.confirmed_date).toLocaleDateString()}
            </p>
          </CardDescription>
        </CardContent>
      ) : (
        <CardContent className="flex-1 h-20 max-h-20">
          <div className="space-y-1">
            <p className="text-sm font-medium leading-none">Proposed Dates:</p>
            <ul className="list-disc pl-5">
              {event.proposed_dates.map((date, index) => (
                <li key={index} className="text-sm text-muted-foreground">
                  {new Date(date).toLocaleDateString()}
                </li>
              ))}
            </ul>
          </div>
        </CardContent>
      )}
      <CardContent>
        <CardDescription>Location: {event.location}</CardDescription>
      </CardContent>
      <CardContent>
        <CardDescription
          className={cn({
            "text-red-500": event.status === "rejected",
            "text-green-500": event.status === "confirmed",
            "text-orange-500": event.status === "pending",
          })}
        >
          Status: {event.status}
        </CardDescription>
      </CardContent>
      <CardContent>
        <CardDescription>
          Date created: {new Date(event.created_at).toLocaleDateString()}
        </CardDescription>
      </CardContent>
      <CardFooter>
        <PopUpViewEvent event={event} role={props.role} />
      </CardFooter>
    </Card>
  );
}

export default EventCard;
