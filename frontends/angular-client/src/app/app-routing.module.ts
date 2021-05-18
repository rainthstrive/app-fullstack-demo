import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LanglistComponent } from './langlist/langlist.component';
import { LangdetailsComponent } from './langdetails/langdetails.component'
const routes: Routes = [
  { path: '', component: LanglistComponent },
  { path: 'details/:id', component: LangdetailsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }