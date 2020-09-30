import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { SocketService } from './socket.service';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TopdriveStreamService } from './webSocketStream.service';

@NgModule({
  declarations: [AppComponent],
  imports: [BrowserModule, AppRoutingModule, FormsModule, HttpClientModule],
  providers: [SocketService, TopdriveStreamService],
  bootstrap: [AppComponent]
})
export class AppModule {}
