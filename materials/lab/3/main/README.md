I chose option two and expanded host.go to take more complex queries.
The terminal gives instructions on how to use it. You will be asked how many filters you want, then you can enter in the filters in the format described. Then, you will be asked to provide the number of pages you would like to view. The terminal instructions read:

-------------------------Welcome to Shodan, Terminal Edition-------------------------
You will be asked go supply filters to search by. The available filters are
os:
timestamp:
isp:
asn:
hostnames:
location:
ip:
domains:
org:
data:
port:
ipstring

Furthermore, location can have a few different ways to search by. These include:
city
region_code:
area_code:
longitude:
country_code3:
country_name:
postal_code:
dma_code:
country_code:
latitude:

Now notice the colon after every filter. You'll need to make sure you have those. An example:
city:Chicago or port:80
------------------------------------------------------------------------
If you have any questions, please contact support at wbrant@uwyo.edu
------------------------------------------------------------------------