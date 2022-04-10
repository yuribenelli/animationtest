
import { Component, OnInit } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { Character } from 'src/app/model/character.model';

@Component({
  selector: 'app-charlist',
  templateUrl: './charlist.component.html',
  styleUrls: ['./charlist.component.scss']
})
export class CharlistComponent implements OnInit {
  charactersList : Character[] = [];
  private http: HttpClient;
  constructor(h:HttpClient) {
    this.http = h;
   }

  ngOnInit(): void {
    this.refresh();
  }
  refresh():void{
    this.http.get<Character[]>(environment.apiUrl + "character").subscribe(data => this.populate(data));
  }
  private populate(data: Character[]){
    this.charactersList = data;
  }

}
