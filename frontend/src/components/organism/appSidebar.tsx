import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
} from "@/components/ui/sidebar";
import { Button } from "../ui/button";
import { DialogCreateEvent } from "./popUpCreateEvent";
import useStore from "@/store/useStore";


export function AppSidebar(prop: { role: string }) {
  const { reset } = useStore();

  return (
    <Sidebar>
      <SidebarHeader />
      <SidebarContent>
        <SidebarGroup>
          {prop.role === "hr" ? <DialogCreateEvent /> : null}
        </SidebarGroup>
        <SidebarGroup />
      </SidebarContent>
      <SidebarFooter>
        <Button onClick={reset}>Logout</Button>
      </SidebarFooter>
    </Sidebar>
  );
}
