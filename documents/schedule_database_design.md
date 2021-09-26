#DatabaseDesign

#### User story

- User can book within future 1 weeks.
- Doctor can set the special date time.

#### Solution

- Week Schedule
  
  - enable_status
  - week_name
  - start_time_am
  - end_time_am
  - start_time_pm
  - end_time_pm
  - number_per_slot
  - appointment_type **[office || virtual]**
  - delete

- Close Date Time
  
  - date
  - start_time_am
  - end_time_am
  - start_time_pm
  - end_time_pm
  - is_disable_this_whole_day