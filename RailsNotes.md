# Rails Notes

[Passing variables to Javascript](#passing-variables-to-javascript)<br/>
[Accessing nested parameter attributes](#accessing-nested-parameter-attributes)<br/>
[Solutions to Errors](#solutions-to-errors)<br/>
&nbsp;&nbsp;&nbsp;[simple_form, passing instance variables](#simple-form,-passing-instance-variables)<br/>

## Passing variables to Javascript

Controller
```ruby
@number = 100
```
View
```javascript
<script>
 var number = <%= @number %>;
</script>
```
## Accessing nested parameter attributes

```ruby
item_params = {
  "name" => "",
  "description" => "",
  "tag_list" => "",
  "user_items_attributes" => {"0" => {"user_id" => "8"}},
  "created_by" => 8,
  "status" => 0
}
```

```ruby
params[:albums][:phots]['0']
```


## Solutions to Errors

### simple_form, passing instance variables.
routes.rb
```ruby
resources :albums do
  resources :photos
end
```
This will create,
```
album_photos_path
```
Intuitively, simple form should look like this,
having two instance variables, because the path has two
parameters.

```rails
<%= simple_form_for([@album, @photo]) do |f| %>
  <%= f.input :title %>
  <%= f.input :description %>
  <%= f.button :submit %>
<% end %>
```
but it will prompt this error,
```
undefined method photos_path' for #<#<Class:0x5232b40>:0x3cd55a0>
```

Solution:
```rails
<%= simple_form_for @photo, url: album_photos_path do |f| %>
  <%= f.input :title %>
  <%= f.input :description %>
  <%= f.button :submit %>
<% end %>
```
or instead keep the previous simple form then,

Controller:
```ruby
def new
  @album = Album.new
  @photo = @album.photos.build
end
```





