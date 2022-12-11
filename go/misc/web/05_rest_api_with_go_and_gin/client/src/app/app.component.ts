import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface INewsfeedItem {
  title: string
  post: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  public title = ''
  public post = ''
  public newsfeedItems: INewsfeedItem[] = []

  constructor (
    private httpClient: HttpClient
  ) {}

  async ngOnInit() {
    await this.loadNewsItems()
  }

  async loadNewsItems() {
    this.httpClient
      .get<INewsfeedItem[]>("/api/newsfeed").subscribe(
        (response) => { this.newsfeedItems = response; }
      )
  }

  async addPost() {
    this.httpClient.post('/api/newsfeed', {
      title: this.title,
      post: this.post,
    }).subscribe(
      (response) => { this.loadNewsItems(); }
    )

    this.title = this.post = ''
  }
}
