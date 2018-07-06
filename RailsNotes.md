# Rails Notes

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
