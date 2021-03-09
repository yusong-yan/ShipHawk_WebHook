require 'sinatra'
require 'pp'

configure { set :server, :puma 
            set :lock, true          
}

$request_number = 0

get '/' do
  "Hello World!"
end

post '/' do
  $request_number= $request_number+1
  request_payload = JSON.parse(request.body.read, symbolize_names: true)
  event_number = request_payload[:totalEvents]
  events = request_payload[:eventList] 
  pp "Request number: #{$request_number}#########################################################"
  #get each event
  events.each do |event|
    #JSON FORMAT FOR EACH EVENT
    #event => { 
    #           trackingNumber 
    #           carrier
    #           estimatedDeliveryDate
    #           scanDetails =>{
    #             eventDate
    #             eventTime
    #             eventCity
    #             eventStateOrProvince
    #             postalCode
    #             country
    #             scanType
    #             scanDescription
    #             packageStatus
    #           }
    #         }
   trackingNumber = event[:trackingNumber]
   carrier = event[:carrier]
   estimatedDeliveryDate = event[:estimatedDeliveryDate]
   scanDetails = event[:scanDetails]
   eventDate = scanDetails[:eventDate]
   eventTime = scanDetails[:eventTime]
   eventCity = scanDetails[:eventCity]
   eventStateOrProvince = scanDetails[:eventStateOrProvince]
   postalCode = scanDetails[:postalCode]
   country = scanDetails[:country]
   scanType = scanDetails[:scanType]
   scanDescription = scanDetails[:scanDescription]
   packageStatus = scanDetails[:packageStatus]

   pp "UPDATE for item #{trackingNumber} =>"
   pp "   Carrier: #{carrier}"
   pp "   EstimatedDeliveryDate: #{estimatedDeliveryDate}"
   pp "   ScanDetails =>"
   pp "       EventDate : #{eventDate}"
   pp "       EventTime : #{eventTime}"
   pp "       EventCity : #{eventCity}"
   pp "       EventStateOrProvince : #{eventStateOrProvince}" 
   pp "       PostalCode : #{postalCode}"
   pp "       Country : #{country}"
   pp "       ScanType : #{scanType}"
   pp "       scanDescription : #{scanDescription}"
   pp "       packageStatus : #{packageStatus}"
   pp ""
  end

  pp "########################################################################"
end
