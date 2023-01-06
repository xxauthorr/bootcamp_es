# BootCampers Esports
Hi Mates, This is my first restful api project done using Go Language. It took me over three weeks to complete the project and i have included multiple api's and packages in it.The whole project is done on popular golang framework called Gin. 

The project is dedicated to the esports community. It is basically a platform for gamers and their associates. Through this application anyone can able to host or book their tournaments and practice matches for any type of e-games, User can create team and add his gaming mates and start a new organization, Players or the team leader can able to send team join requests and the teams even able to add team recruitment notices. These things can be done after creating an account in BootCampers Esports. It is basically an alternative of discord but in BootCampers Esports it is much more easy to make it happen because here i only included the things that are really needed for the gamers. 

It is a very secured application that i have used the JWT authorization and OTP signup. And i tried to bring a lot of features in this release and i have a bundle of other features that i will be including in the future. I have completed this project in very good coding standards and it is flexible, So any programers can easily understand the code and easily adapt the working flow of the project. I have followed the MVC code architecture which is very important while building a large scale program. These all made the project flexible and scalable.

## Deployment 

So as it's my first project i done hosting in basic methods, So i can cover all the deployment methods from the foundation which will help me understand each steps in future when am working on new techonologies like docker, kubernetes, and all those deployment tools. 

I Deployed my project in AWS cloud service called EC2 which is a remote instance(server). In that i used the nginx web server for handling the reverse proxy and configuring the web application to secure HTTPS.

## Database

I have used the RDS service which is cloud relational database service provided by the AWS(Amazon web service) in which we can deploy or use any type of SQL databases. I have done the project using the PSQL(Postgre SQL) which is free to use until certain amount of usage.

## File Server

For storing the files from the client i used the AMAZON S3 bucket which is also very popular service provided by the amazon. And i created an IAM role to connect this bucket with the instance and to the program.

## API's Used 

This project is much more scalable than this stage but still i have included multiple api's to this project which are listed below.

Twilio   - Which helps with the OTP authentication 

Youtube  - As it is an esports platform i bring the top game streaming live to this page which would encourage the gamers to get promoted and notified.

AmazonS3 - This is used to upload the files to the amazon file server.

## Future Add on's

-Will be adding web sockets to bring chat within the players and in the team.

-More api's like youtube to showcase the players skills

-Will be including merchandise shop for the teams where they can sell their brand products

-Will be including pubg,clash of clans,call of duty,fortnite,etc.. api's to verify the user and bring their gaming experience status 

-Live streaming gameplays can also may be added in the future



