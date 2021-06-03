using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Services;
using ApiFarm.Utils;
using ApiFarm.Utils.Impl;
using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Controllers
{
    /// <summary>
    /// Controller for consuming requests to get, add, update or delete video game entities from storage.
    /// </summary>
    [NotSupportedExceptionFilter]
    [Route("video_games")]
    [ApiController]
    public class VideoGamesController : ControllerBase
    {
        private IService<VideoGame> videoGameService;

        /// <summary>
        /// Initializes a new instance of the <see cref="VideoGamesController"/> class.
        /// </summary>
        /// <param name="videoGameService">Service to query for creating, reading, updating and deleting <see cref="VideoGame"/> resources.</param>
        public VideoGamesController(IService<VideoGame> videoGameService)
        {
            this.videoGameService = videoGameService;
        }

        /// <summary>
        /// Endpoint to retrieve <see cref="VideoGame"/> model from storage.
        /// </summary>
        /// <param name="strId">Identifier for <see cref="VideoGame"/> to retrieve as a string.</param>
        /// <returns>Successful result with queried <see cref="VideoGame"/> or a Not Found response.</returns>
        [HttpGet]
        [Route("{strId}")]
        public ObjectResult Get(string strId)
        {
            if (!uint.TryParse(strId, out var id))
            {
                return this.BadRequest(ResponseMessages.Id.IsInvalid(strId));
            }

            var query = this.videoGameService.Get(id);

            if (query.Code != 0)
            {
                return this.NotFound(query.Message);
            }

            return this.Ok(query.Result);
        }

        /// <summary>
        /// Endpoint to retrieve all <see cref="VideoGame"/> models from storage.
        /// </summary>
        /// <returns>All video games in an array as JSON.</returns>
        [HttpGet]
        public ActionResult<IEnumerable<VideoGame>> GetAll()
        {
            var query = this.videoGameService.GetAll();
            return this.Ok(query.Result);
        }

        /// <summary>
        /// Endpoint to add a <see cref="VideoGame"/> to storage.
        /// </summary>
        /// <param name="videoGame">The <see cref="VideoGame"/> extracted from request body to be added to storage.</param>
        /// <returns>The video game which was added.</returns>
        [HttpPost]
        [JsonResourceFilter]
        public ObjectResult Post(VideoGame videoGame)
        {
            var query = this.videoGameService.Add(videoGame);

            if (query.Code != 0)
            {
                return this.BadRequest(query.Message);
            }

            videoGame = query.Result;
            return this.Created($"video_games/{videoGame.Id}", videoGame);
        }

        /// <summary>
        /// Endpoint to update a <see cref="VideoGame"/> in storage.
        /// </summary>
        /// <param name="strId">Identifier of <see cref="VideoGame"/> to update in storage as string.</param>
        /// <param name="videoGameUpdateValues"><see cref="VideoGame"/> with set fields to be updated in storage.</param>
        /// <returns>Successful query with updated <see cref="VideoGame"/>.</returns>
        [HttpPut]
        [Route("{strId}")]
        [JsonResourceFilter]
        public ObjectResult Put(string strId, VideoGame videoGameUpdateValues)
        {
            if (!uint.TryParse(strId, out var id))
            {
                return this.BadRequest(ResponseMessages.Id.IsInvalid(strId));
            }

            var query = this.videoGameService.Update(id, videoGameUpdateValues);

            if (query.Code != 0)
            {
                return this.NotFound(query.Message);
            }

            return this.Ok(query.Result);
        }
    }
}
