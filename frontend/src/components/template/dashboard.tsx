/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import { AppSidebar } from '../organism/appSidebar';
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import GridContainer from './gridContainer';
import { useEvent } from '@/hooks/useEvent';



const HrDashboard = (prop: {role: string}) => {
    const { events } = useEvent();

    return (
        <SidebarProvider>
          <AppSidebar role={prop.role}/>
          <main>
            <SidebarTrigger />
            {
              events!= undefined ? (
                <GridContainer events={events} role={prop.role}/>
              ) : (
                <div className="flex items-center justify-center h-full">
                  <h1 className="text-2xl text-gray-500">No event found</h1>
                </div>
              )
            }
          </main>
        </SidebarProvider>
      )
};

export default HrDashboard;