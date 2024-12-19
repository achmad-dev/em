import React from 'react';
import type { Event } from '@/type/type';
import EventCard from '../organism/eventCard';

interface GridContainerProps {
    events: Event[];
    role: string;
}

const GridContainer: React.FC<GridContainerProps> = ({ events, role }) => {
    return (
        <div className="w-full bg-white rounded-lg shadow-sm p-4 flex flex-wrap gap-4">
            {events.length > 0 ? events
            .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
            .map((event, index) => (
                <div key={index} className="w-full sm:w-[calc(50%-0.5rem)] md:w-[calc(33.33%-0.67rem)] lg:w-[calc(25%-0.75rem)] max-w-[300px]">
                <EventCard key={index} event={event} role={role} />
                </div>
            )) : null}
        </div>
    );
};

export default GridContainer;