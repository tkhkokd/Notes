# Practical Ruby

## Running a code periodically

```ruby
  def every_n_seconds(n)
    loop do
      before = Time.now
      yield
      interval = n-(Time.now-before)
      sleep(interval) if interval > 0
    end
  end

  every_n_seconds(5) do
   puts "#{Time.now.strftime("%X")}"
  end
```

## Joining two sets of data and create a master data

```ruby
require "csv"

# Load the CSV files, store them in different variables.
# Each file will be imported as a single string, with \r\n as the delimeter for the rows and "," for the columns.
bundles = File.read("./bundles.csv")
bundle_items = File.read("./bundle_items.csv")

# Separate each string into an array of strings(rows),
# then convert each string(a row) into an array of bundle_id and xxx_SKU,
# Drop the header.
bundles = bundles.split(/\r\n/).map{|b| b.split(',')}.drop(1)
bundle_items = bundle_items.split(/\r\n/).map{|bi| bi.split(',')}.drop(1)

# Initiate variables for the master data as hashes
# Hash structure:
# Bundles master data: {bundle_id => bundle_SKU}, Bundle items master data: {bundle_id => [item_SKU item_SKU]}
master_bundles = {}
master_bundle_items = Hash.new{|hsh, key| hsh[key] = []}

# Push data into the hashes
bundles.map{|b| master_bundles[b[0]] = b[1]}
bundle_items.map{|bi| master_bundle_items[bi[0]] << bi[1]}

# Initiate the desired new table data.
new_table = []
# Iterate the bundles hash and get bundle_SKU(v),
# and use its key(k) to get item_SKU and store them in new_table.
# Concatenate the item_SKUs with spaces.
master_bundles.each {|k, v| new_table << [v, master_bundle_items[k].join(" ")]}

# Create a new csv file with the specified header,
# insert each row(array) from new_table to the new csv.
CSV.open("./new_table.csv", "wb", :write_headers => true, :headers => ["bundle_SKU", "item_SKU"]) do |csv|
  new_table.each{|row| csv << row }
end
```
