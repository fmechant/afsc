## Feb 25 2019 0.14.6
  * Added gs.DefaultHTTPClientProvider
  * Added gs.DefaultProjectProvider
  
## Feb 7 2019 0.14.2
  * Added preSing option on s3 (get, upload)
  * Added upload content meta support option

## December 20 2019 0.12.0
  * Added Google Cloud Storage http proxy support
  * Refactored retries 
  
## December 1 2019 0.10.8
  * Patched object kind support  
   
## November 19 2019 0.9.1
  * Optimize Exists, Get operation (to avoid expensive list operation)
  
## November 5 2019 0.7.4
  * Added upload storage.Object option (to reduce Object call)
  
## November 5 2019 0.7.3
  * Added retry on backenError
  * Added sleep retry on backendError
  
## November 1 2019 0.7.0
 * Added streaming option for move/copy fallback
 * Added s3 logging
 * Renamed Checksum to SkipChecksum reference
 
## October 28 2019 0.6.1
  * Patched initial over- allocation on s3 reader
  
## October 27 2019 0.6.0
  * Implemented Stream Reader
  * Added support for Steam/Checksum option
  * Update Storage signature changes (breaking changes)

## October 25 2019 0.4.1
  * Implemented AuthTracker
  
## October 23 2019 0.3.4
  * Patched error shadowing on move error fallback
  * Added region bucket loading 

## October 17 2019 0.3.2
  * Added move fallback on google storage fallback in case storage class incompatibility
  
## October 15 2019 0.3.1
  * Implemented copier interface
  * Added copy fallback on google storage fallback in case storage class incompatibility  

## October 2 2019 0.2.2
  * Patched s3 exists
  * Updated s3 storager close
  * Added tests

## October 1 2019 0.2.0

  * Patched bucket creation
  * Updated matcher handling
  
  
  
## Agust 20 2018

  * Initial Release.


