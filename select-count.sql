SELECT 'airport', COUNT(*) from airport
UNION ALL
SELECT 'attendant', COUNT(*) from attendant
UNION ALL
SELECT 'booking_office', COUNT(*) from booking_office
UNION ALL
SELECT 'cashier', COUNT(*) from cashier
UNION ALL
SELECT 'line', COUNT(*) from line
UNION ALL
SELECT 'plane_model', COUNT(*) from plane_model
UNION ALL
SELECT 'plane', COUNT(*) from plane
UNION ALL
SELECT 'pilot', COUNT(*) from pilot
UNION ALL
SELECT 'flight', COUNT(*) from flight
UNION ALL
SELECT 'seat', COUNT(*) from seat
UNION ALL
SELECT 'tariff', COUNT(*) from tariff
UNION ALL
SELECT 'purchase', COUNT(*) from purchase
UNION ALL
SELECT 'passenger', COUNT(*) from passenger
UNION ALL
SELECT 'ticket', COUNT(*) from ticket
UNION ALL
SELECT 'flight_in_ticket', COUNT(*) from flight_in_ticket
UNION ALL
SELECT 'pilot_flies_plane_model', COUNT(*) from pilot_flies_plane_model
UNION ALL
SELECT 'flight_attendant', COUNT(*) from flight_attendant;