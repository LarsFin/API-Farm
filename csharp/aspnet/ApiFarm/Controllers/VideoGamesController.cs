using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Controllers
{
    [Route("video_games")]
    [ApiController]
    public class VideoGamesController : ControllerBase
    {
        [HttpGet]
        public ActionResult<string> Get()
        {
            return "Hello World!";
        }
    }
}
